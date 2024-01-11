package asr

import (
	pb "applet-server/internal/biz/asr/proto"
	"applet-server/internal/conf"
	"applet-server/internal/data"
	"applet-server/internal/pkg/log"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"strconv"
	"time"
)

type AsRControllerClient struct {
	*grpc.ClientConn
	*log.MyLogger
}

func NewAsRControllerClient(c *conf.App, logger *log.MyLogger) *AsRControllerClient {
	ctx, _ := context.WithTimeout(context.Background(), c.Asr.GetTimeout().AsDuration())
	conn, err := grpc.DialContext(ctx, c.Asr.GetAddr(),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	return &AsRControllerClient{
		ClientConn: conn,
		MyLogger:   logger,
	}
}

func (c *AsRControllerClient) StreamingRecognize(ctx context.Context, session *data.Session, voiceDataCh chan []byte, asrRecognizedText chan string) error {
	client := pb.NewSpeechClient(c.ClientConn)
	streamClient, err := client.StreamingRecognize(ctx)
	if err != nil {
		c.Infof("streamingRecognize error: %v", err)
		return err
	}

	// 接收流式返回结果
	go ReceiveRecognizedText(streamClient, asrRecognizedText)

	recognitionRequest := newRecognitionRequest(strconv.Itoa(int(session.RobotId)), session.Id, session.Language.Load(), session.TraceId, session.AgentId)
	awaitTime := 30 * time.Second
	vadTimer := time.NewTimer(awaitTime)
	c.Debug("recognitionRequest:", recognitionRequest)

	for {
		select {
		case voiceData, ok := <-voiceDataCh:
			if ok {
				c.Debug("StreamingRecognize voice data; the length", len(voiceData))
				recognitionRequest.Body.Data.Speech = voiceData
				err = streamClient.Send(recognitionRequest)
				// 重置超时定时器
				vadTimer.Reset(awaitTime)
				if err != nil {
					return err
				}
			} else {
				goto END
			}

		case <-vadTimer.C:
			c.Debug("StreamingRecognize timeout")
			goto END
		}
	}
END:
	recognitionRequest.Body.Data.Speech = nil
	recognitionRequest.Extra = &pb.Extra{ExtraType: "audioExtra", ExtraBody: "val"}

	if err = streamClient.Send(recognitionRequest); err != nil {
		return err
	}
	streamClient.CloseSend()
	return nil
}

func ReceiveRecognizedText(streamClient pb.Speech_StreamingRecognizeClient, recognizedTextCh chan string) {
	defer close(recognizedTextCh)
	for {
		response, err := streamClient.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		if response.DetailMessage != nil {
			recognizedText := (response.DetailMessage.Fields["recognizedText"]).GetStringValue()
			if recognizedText != "" {
				recognizedTextCh <- recognizedText
			}
		}
	}

}

func newRecognitionRequest(robotId, sessionId, language, traceId string, agentId int) *pb.RecognitionRequest {
	vendor := "CloudMinds"
	dialect := "zh"
	if language == "EN" {
		vendor = "Google"
		dialect = "en-US"
	}
	recognitionRequest := &pb.RecognitionRequest{
		CommonReqInfo: &pb.CommonReqInfo{
			Guid:        fmt.Sprintf("sv_controller_%s", sessionId),
			TenantId:    "jshuibo",
			RobotId:     robotId,
			UserId:      "wechat_applets",
			Version:     "1.0",
			ServiceCode: "VirtualRobot",
			Timestamp:   time.Now().Unix(),
			RootGuid:    traceId,
		},
		Body: &pb.Body{
			Option: map[string]string{
				"returnDetail":  "true",
				"returnTrace":   "true",
				"recognizeOnly": "true",
				"tstAgentId":    strconv.Itoa(agentId),
				"tstAsrDomain":  "Common_V2",
			},
			Data: &pb.Body_Data{
				Rate:     16000,    // means sampling-rate always 16000
				Format:   "pcm",    // only pcm
				Language: language, // CH EN TCH JA ES
				Dialect:  dialect,  //例（百度：1537，讯飞：zh_cn，谷歌：zh）
				// when vendor=Google then "https://cloud.google.com/speech-to-text/docs/languages" for column "languageCode" (例zh)
				// when vendor=Baidu then "https://ai.baidu.com/ai-doc/SPEECH/ek38lxj1u#%E8%AF%86%E5%88%AB%E6%A8%A1%E5%9E%8Bdev_pid%E5%8F%82%E6%95%B0" for section "dev_pid 参数列表" (例1537)
				// when vendor=IFlyTek then "https://www.xfyun.cn/doc/asr/voicedictation/API.html#%E6%8E%A5%E5%8F%A3%E8%B0%83%E7%94%A8%E6%B5%81%E7%A8%8B" for section "业务参数" "language" (例zh_cn)
				Vendor: vendor, //Baidu,IFlyTek,IFlyTekRealTime,Google,CloudMinds,inactiveAsrClient,CPALL
			},
			Type: 1,
		},
	}
	return recognitionRequest
}
