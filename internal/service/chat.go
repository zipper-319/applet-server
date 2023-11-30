package service

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/biz/asr"
	"applet-server/internal/biz/nlp"
	"applet-server/internal/biz/tts"
	"applet-server/internal/data"
	"applet-server/internal/pkg/ws"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"strings"
	"sync"
)

type ChatService struct {
	log.Logger
	*tts.TTSClient
	*nlp.TalkClient
	*asr.AsRControllerClient
}

func NewChatService(helper log.Logger, ttsClient *tts.TTSClient, asrClient *asr.AsRControllerClient, talkClient *nlp.TalkClient) *ChatService {
	return &ChatService{
		Logger:              helper,
		TTSClient:           ttsClient,
		TalkClient:          talkClient,
		AsRControllerClient: asrClient,
	}
}

func (c *ChatService) HandlerVoice(ctx context.Context, vadOutChan chan []byte, conn *websocket.Conn, session *data.Session) error {
	ctx, cancel := context.WithCancel(ctx)
	logger := log.NewHelper(log.With(c.Logger, "sessionId", session.Id, "traceId", session.TraceId))
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("HandlerVoice; err:%v", err)
		}
		cancel()
		logger.Debugf("traceId:%s; end to  handler Voice; session:%v", session.TraceId, session)
	}()
	logger.Debugf("traceId:%s; start to  handler Voice; session:%v", session.TraceId, session)

	asrOutChan := make(chan string, 10)
	if err := c.AsRControllerClient.StreamingRecognize(ctx, session, vadOutChan, asrOutChan); err != nil {
		return err
	}
	asrResult := ""
	logger.Debugf("traceId:%s; await to asrResult ", session.TraceId)
	for asrText := range asrOutChan {
		asrResult = asrText
		logger.Debugf("traceId:%s, asrResult:%s", session.TraceId, asrResult)
		if asrResult != "" {
			ws.SendingMsgToClient(conn, applet.ServiceType_Service_ASR, asrText)
		}
	}
	ws.SendFinishedMsgToClient(conn, applet.ServiceType_Service_ASR, asrResult)

	talkRespCh := make(chan data.TalkResp, 10)
	if err := c.TalkClient.StreamingTalkByText(ctx, asrResult, session, talkRespCh); err != nil {
		return fmt.Errorf("talk by text error: %s", err.Error())
	}

	var wg sync.WaitGroup

	for resp := range talkRespCh {
		logger.Debugf("traceId:%s, resp:%v", session.TraceId, resp)
		if err := ws.SendingMsgToClient(conn, applet.ServiceType_Service_Nlp, resp); err != nil {
			return err
		}
		for _, answer := range resp.AnsItem {
			ttsText := strings.Replace(strings.TrimSpace(answer.Text), "\n", "", -1)
			logger.Debugf("sessionId:%s, ttsText:%s; start to call tts v2", session.TraceId, ttsText)
			if ttsText == "" {
				continue
			}
			wg.Add(1)
			go c.HandlerTTSToClient(ctx, logger, ttsText, answer.Lang, session, conn, &wg)

		}
	}
	if err := ws.SendFinishedMsgToClient(conn, applet.ServiceType_Service_Nlp, ""); err != nil {
		return err
	}
	logger.Debugf("sessionId:%s, the sentence finished", session.Id)
	return nil
}
func (c *ChatService) HandlerText(ctx context.Context, body string, conn *websocket.Conn, session *data.Session) error {
	logger := log.NewHelper(log.With(c.Logger, "sessionId", session.Id, "traceId", session.TraceId, "method", session.MethodType))
	logger.Infof("text:%s", body)
	if session == nil {
		return errors.New("session is nil")
	}
	var wg sync.WaitGroup
	if session.MethodType == applet.MethodType_OnlyTTS {

		wg.Add(1)
		go c.HandlerTTSToClient(ctx, logger, body, "", session, conn, &wg)

	} else {
		talkRespCh := make(chan data.TalkResp, 10)
		if err := c.StreamingTalkByText(ctx, body, session, talkRespCh); err != nil {
			return fmt.Errorf("talk by text error: %s", err.Error())
		}

		for resp := range talkRespCh {

			if err := ws.SendingMsgToClient(conn, applet.ServiceType_Service_Nlp, resp); err != nil {
				return err
			}
			for _, answer := range resp.AnsItem {
				ttsText := strings.Replace(strings.TrimSpace(answer.Text), "\n", "", -1)
				log.Debugf("sessionId:%s, ttsText:%s; start to call tts v2", session.TraceId, ttsText)
				if ttsText == "" {
					continue
				}
				wg.Add(1)
				go c.HandlerTTSToClient(ctx, logger, ttsText, answer.Lang, session, conn, &wg)
			}
		}
	}

	if err := ws.SendFinishedMsgToClient(conn, applet.ServiceType_Service_Nlp, ""); err != nil {
		return err
	}
	wg.Wait()

	logger.Debugf("HandlerText finished;")
	return nil
}

func (c *ChatService) HandlerTTSToClient(ctx context.Context, logger *log.Helper, ttsText, language string, session *data.Session, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer func() {
		logger.Infof("sessionId:%s,ttsText:%s, tts finished", session.TraceId, ttsText)
		wg.Done()
	}()
	ttsParam := session.TtsParam.Load().(*data.TTSParam)
	log.Debugf("start to call tts; sessionId:%s, ttsText:%s, ttsParam:%+v", session.TraceId, ttsText, ttsParam)
	ttsChan, err := c.CallTTSV2(ctx, ttsParam, ttsText, language, session.Id, session.TraceId)
	if err != nil {
		logger.Errorf("sessionId:%s,ttsText:%s, call tts error:%v", session.TraceId, ttsText, err)
		return
	}
	for ttsResp := range ttsChan {
		if err := ws.SendingMsgToClient(conn, applet.ServiceType_Service_TTS, base64.RawStdEncoding.EncodeToString(ttsResp)); err != nil {
			logger.Errorf("sessionId:%s,ttsText:%s, send tts error:%v", session.TraceId, ttsText, err)
		}
	}
}
