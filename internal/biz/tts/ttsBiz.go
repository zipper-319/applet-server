package tts

import (
	"applet-server/internal/biz/tts/proto/v1"
	"applet-server/internal/biz/tts/proto/v2"
	"applet-server/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
)

type TTSClient struct {
	*grpc.ClientConn
	*log.Helper
}

func NewTTSClient(c *conf.App, logger log.Logger) *TTSClient {
	ctx, _ := context.WithTimeout(context.Background(), c.Tts.GetTimeout().AsDuration())
	conn, err := grpc.DialContext(ctx, c.Tts.GetAddr(),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	return &TTSClient{
		ClientConn: conn,
		Helper:     log.NewHelper(log.With(logger, "Model", "TTS")),
	}
}

func (c *TTSClient) CallTTSV2(ctx context.Context, speaker, text, sessionId string, robotId int32) (chan []byte, error) {
	traceId := uuid.New().String()
	ttsV2Client := v2.NewCloudMindsTTSClient(c.ClientConn)
	req := &v2.TtsReq{
		Text:                 text,
		ParameterSpeakerName: speaker,
		TraceId:              traceId,
		RootTraceId:          sessionId,
		Version:              v2.ClientVersion_version,
	}

	response, err := ttsV2Client.Call(ctx, req)
	if err != nil {
		log.Errorf("traceId:%s,Text:%s, err;%v", traceId, text, err)
		return nil, err
	}

	ttsChan := make(chan []byte, 10)

	go func() {
		defer func() {
			log.Debugf("sessionId:%s, traceId:%s; text:%s,finish to CallTTSV2, next to send tts video", sessionId, traceId, text)
			close(ttsChan)
		}()
		var number = 0
		for {
			select {
			case <-ctx.Done():
				log.Infof("traceId:%s;  TestTTSV2 return after cancel\n", traceId)
				return
			default:
				temp, err := response.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Errorf("sessionId:%s,traceId:%s,Text:%s, err;%v", sessionId, traceId, text, err)
					return
				}
				if temp.ErrorCode != 0 {
					log.Errorf("tts 内部服务错误：%d", temp.ErrorCode)
				}

				if audio, ok := temp.ResultOneof.(*v2.TtsRes_SynthesizedAudio); ok {
					number += 1
					log.Infof("index:%d, pcm length:%d, status:%d", number, len(audio.SynthesizedAudio.Pcm), temp.Status)
					ttsChan <- audio.SynthesizedAudio.Pcm
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
		log.Errorf("traceId:%s,Text:%s, err;%v", traceId, text, err)
		return nil, err
	}

	ttsChan := make(chan []byte, 10)

	go func() {
		defer func() {
			log.Debugf("sessionId:%s, traceId:%s; text:%s,finish to CallTTSV2, next to send tts video", sessionId, traceId, text)
			close(ttsChan)
		}()
		var number = 0
		for {
			select {
			case <-ctx.Done():
				log.Infof("traceId:%s;  TestTTSV2 return after cancel\n", traceId)
				return
			default:
				temp, err := response.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Errorf("sessionId:%s,traceId:%s,Text:%s, err;%v", sessionId, traceId, text, err)
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
