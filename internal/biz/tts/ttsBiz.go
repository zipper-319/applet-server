package tts

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/biz/tts/proto/v1"
	"applet-server/internal/biz/tts/proto/v2"
	"applet-server/internal/conf"
	"applet-server/internal/data"
	"applet-server/internal/pkg/log"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
	"time"
)

type TTSClient struct {
	*grpc.ClientConn
	*log.MyLogger
	timeout time.Duration
	*data.Data
}

func NewTTSClient(c *conf.App, data *data.Data, logger *log.MyLogger) *TTSClient {
	timeout := c.Tts.GetTimeout().AsDuration()
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	conn, err := grpc.DialContext(ctx, c.Tts.GetAddr(),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	return &TTSClient{
		ClientConn: conn,
		MyLogger:   logger,
		timeout:    timeout,
		Data:       data,
	}
}

func (c *TTSClient) GetSpeechClient(env applet.EnvType) (v2.CloudMindsTTSClient, error) {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout)
	conn, err := grpc.DialContext(ctx, c.GetTTSAddr(string(env)),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	client := v2.NewCloudMindsTTSClient(conn)
	return client, nil
}

func (c *TTSClient) CallTTSV2(ctx context.Context, session *data.Session, ttsParam *data.TTSParam, text, language string) (chan []byte, error) {
	username := session.Username
	traceId := ctx.Value("questionId").(string)
	sessionId := ctx.Value("sessionId").(string)
	ttsV2Client, err := c.GetSpeechClient(session.Env)
	if err != nil {
		return nil, err
	}
	req := &v2.TtsReq{
		Text:                 text,
		ParameterSpeakerName: ttsParam.Speaker,
		Volume:               ttsParam.Volume,
		Speed:                ttsParam.Speed,
		Pitch:                ttsParam.Pitch,
		TraceId:              traceId,
		RootTraceId:          sessionId,
		Language:             language,
		Version:              v2.ClientVersion_version,
	}

	if ttsParam.IsClone {
		req.Userspace = username
	}

	response, err := ttsV2Client.Call(ctx, req)
	if err != nil {
		return nil, err
	}

	ttsChan := make(chan []byte, 10)

	go func() {
		defer func() {
			c.WithContext(ctx).Debugf("sessionId:%s, traceId:%s; text:%s,finish to CallTTSV2, next to send tts video", sessionId, traceId, text)
			close(ttsChan)
		}()
		var number = 0
		var tempAudio []byte
		for {
			select {
			case <-ctx.Done():
				c.Infof("traceId:%s;  TestTTSV2 return after cancel\n", traceId)
				return
			default:
				temp, err := response.Recv()
				if err == io.EOF {
					if len(tempAudio) > 0 {
						ttsChan <- tempAudio
						log.Infof("index:%d, pcm length:%d; end status", number, len(tempAudio))
					}
					return
				}
				if err != nil {
					c.Errorf("sessionId:%s,traceId:%s,Text:%s, err;%v", sessionId, traceId, text, err)
					return
				}
				if temp.ErrorCode != 0 {
					c.Errorf("tts 内部服务错误：%d", temp.ErrorCode)
				}

				if audio, ok := temp.ResultOneof.(*v2.TtsRes_SynthesizedAudio); ok {
					if tempAudio == nil {
						tempAudio = audio.SynthesizedAudio.Pcm
					} else {
						tempAudio = append(tempAudio, audio.SynthesizedAudio.Pcm...)
					}
					number += 1
					log.Infof("index:%d, pcm length:%d, status:%d, tempAudio:%d; IsPunctuation:%d", number, len(audio.SynthesizedAudio.Pcm), temp.Status, len(tempAudio), audio.SynthesizedAudio.IsPunctuation)
					if len(tempAudio) > 0 && audio.SynthesizedAudio.IsPunctuation == 1 {
						ttsChan <- tempAudio
						log.Infof("index:%d, tempAudio length:%d, status:%d", number, len(tempAudio), temp.Status)
						tempAudio = nil
					}
				}
			}
		}
	}()

	return ttsChan, nil
}

func (c *TTSClient) CallTTSV1(ctx context.Context, speaker, text, language, sessionId string, robotId int32) (chan []byte, error) {
	traceId := uuid.New().String()
	ttsV1Client := v1.NewCloudMindsTTSClient(c.ClientConn)
	req := &v1.TtsReq{
		Text:                 text,
		ParameterSpeakerName: speaker,
		TraceId:              traceId,
		RootTraceId:          sessionId,
		Speed:                "3",
		Language:             language,
	}

	response, err := ttsV1Client.Call(ctx, req)
	if err != nil {
		c.WithContext(ctx).Errorf("Text:%s, err;%v", text, err)
		return nil, err
	}

	ttsChan := make(chan []byte, 10)

	go func() {
		defer func() {
			c.WithContext(ctx).Debugf("text:%s,finish to CallTTSV2, next to send tts video", text)
			close(ttsChan)
		}()
		var number = 0
		for {
			select {
			case <-ctx.Done():
				c.WithContext(ctx).Info(" TestTTSV2 return after cancel\n")
				return
			default:
				temp, err := response.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					c.WithContext(ctx).Errorf("Text:%s, err;%v", text, err)
					return
				}
				if temp.Status == v1.PcmStatus_STATUS_MID {
					number += 1
					c.WithContext(ctx).Debugf("index:%d, pcm length:%d, status:%d", number, len(temp.Pcm), temp.Status)
					ttsChan <- temp.Pcm
				}

			}
		}
	}()

	return ttsChan, nil
}

func (c *TTSClient) GetUserSpeaker(ctx context.Context, user string) ([]string, error) {
	traceId := uuid.New().String()
	ttsV2Client := v2.NewCloudMindsTTSClient(c.ClientConn)
	resp, err := ttsV2Client.GetUserSpeakers(ctx, &v2.GetUserSpeakersRequest{
		TraceId: traceId,
		User:    user,
	})
	if err != nil {
		return nil, err
	}
	return resp.Speakers, nil
}

func (c *TTSClient) GetTTSConfig(ctx context.Context, user string) (*v2.RespGetTtsConfig, error) {
	traceId := uuid.New().String()
	ttsV2Client := v2.NewCloudMindsTTSClient(c.ClientConn)
	resp, err := ttsV2Client.GetTtsConfigByUser(ctx, &v2.GetTtsConfigByUserRequest{
		TraceId: traceId,
		User:    user,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
