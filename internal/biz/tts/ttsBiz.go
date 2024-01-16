package tts

import (
	"applet-server/internal/biz/tts/proto/v1"
	"applet-server/internal/biz/tts/proto/v2"
	"applet-server/internal/conf"
	"applet-server/internal/data"
	"applet-server/internal/pkg/log"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
)

type TTSClient struct {
	*grpc.ClientConn
	*log.MyLogger
}

func NewTTSClient(c *conf.App, logger *log.MyLogger) *TTSClient {
	ctx, _ := context.WithTimeout(context.Background(), c.Tts.GetTimeout().AsDuration())
	conn, err := grpc.DialContext(ctx, c.Tts.GetAddr(),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	return &TTSClient{
		ClientConn: conn,
		MyLogger:   logger,
	}
}

func (c *TTSClient) CallTTSV2(ctx context.Context, username string, ttsParam *data.TTSParam, text, language, sessionId, traceId string) (chan []byte, error) {

	ttsV2Client := v2.NewCloudMindsTTSClient(c.ClientConn)
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
		c.Errorf("traceId:%s,Text:%s, err;%v", traceId, text, err)
		return nil, err
	}

	ttsChan := make(chan []byte, 10)

	go func() {
		defer func() {
			c.Debugf("sessionId:%s, traceId:%s; text:%s,finish to CallTTSV2, next to send tts video", sessionId, traceId, text)
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
					log.Infof("index:%d, pcm length:%d, status:%d", number, len(audio.SynthesizedAudio.Pcm), temp.Status)
					if len(tempAudio) > 0 && audio.SynthesizedAudio.IsPunctuation == 1 {
						ttsChan <- tempAudio
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
		c.Errorf("traceId:%s,Text:%s, err;%v", traceId, text, err)
		return nil, err
	}

	ttsChan := make(chan []byte, 10)

	go func() {
		defer func() {
			c.Debugf("sessionId:%s, traceId:%s; text:%s,finish to CallTTSV2, next to send tts video", sessionId, traceId, text)
			close(ttsChan)
		}()
		var number = 0
		for {
			select {
			case <-ctx.Done():
				c.Infof("traceId:%s;  TestTTSV2 return after cancel\n", traceId)
				return
			default:
				temp, err := response.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					c.Errorf("sessionId:%s,traceId:%s,Text:%s, err;%v", sessionId, traceId, text, err)
					return
				}
				if temp.Status == v1.PcmStatus_STATUS_MID {
					number += 1
					log.Infof("index:%d, pcm length:%d, status:%d", number, len(temp.Pcm), temp.Status)
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
