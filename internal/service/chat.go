package service

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/biz/asr"
	"applet-server/internal/biz/nlp"
	"applet-server/internal/biz/tts"
	"applet-server/internal/data"
	"applet-server/internal/pkg/log"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type ChatService struct {
	*log.MyLogger
	*tts.TTSClient
	*nlp.TalkClient
	*asr.AsRControllerClient
}

func NewChatService(ttsClient *tts.TTSClient, asrClient *asr.AsRControllerClient, talkClient *nlp.TalkClient, logger *log.MyLogger) *ChatService {
	return &ChatService{
		MyLogger:            logger,
		TTSClient:           ttsClient,
		TalkClient:          talkClient,
		AsRControllerClient: asrClient,
	}
}

func (c *ChatService) HandlerVoice(ctx context.Context, vadOutChan chan []byte, session *data.Session) error {
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		if err := recover(); err != nil {
			c.Errorf("HandlerVoice; err:%v", err)
		}
		cancel()
		c.Debugf("traceId:%s; end to  handler Voice; session:%v", session.TraceId, session)
	}()
	c.Debugf("traceId:%s; start to  handler Voice; session:%v", session.TraceId, session)

	asrOutChan := make(chan string, 10)
	if err := c.AsRControllerClient.StreamingRecognize(ctx, session, vadOutChan, asrOutChan); err != nil {
		return err
	}
	asrResult := ""
	c.Debugf("traceId:%s; await to asrResult ", session.TraceId)
	for asrText := range asrOutChan {
		asrResult = asrText
		c.Debugf("traceId:%s, asrResult:%s", session.TraceId, asrResult)
		if asrResult != "" {
			session.SendingMsgToClient(applet.ServiceType_Service_ASR, asrText, false, "")
		}
	}
	session.SendingMsgToClient(applet.ServiceType_Service_ASR, asrResult, true, "")

	talkRespCh := make(chan data.TalkResp, 10)
	if err := c.TalkClient.StreamingTalkByText(ctx, asrResult, session, talkRespCh); err != nil {
		return fmt.Errorf("talk by text error: %s", err.Error())
	}

	var wg sync.WaitGroup

	for resp := range talkRespCh {
		c.Debugf("traceId:%s, resp:%v", session.TraceId, resp)
		if err := session.SendingMsgToClient(applet.ServiceType_Service_Nlp, resp, false, ""); err != nil {
			return err
		}
		for _, answer := range resp.AnsItem {
			ttsText := strings.Replace(strings.TrimSpace(answer.Text), "\n", "", -1)
			c.Debugf("sessionId:%s, ttsText:%s; start to call tts v2", session.TraceId, ttsText)
			if ttsText == "" {
				continue
			}
			wg.Add(1)
			go c.HandlerTTSToClient(ctx, ttsText, answer.Lang, session, &wg)

		}
	}
	wg.Wait()
	if err := session.SendingMsgToClient(applet.ServiceType_Service_Nlp, "", true, ""); err != nil {
		return err
	}
	c.Debugf("sessionId:%s, the sentence finished", session.Id)
	return nil
}
func (c *ChatService) HandlerText(ctx context.Context, body string, session *data.Session) error {

	c.Infof("text:%s", body)
	if session == nil {
		return errors.New("session is nil")
	}
	var wg sync.WaitGroup
	if session.MethodType == applet.MethodType_OnlyTTS {

		wg.Add(1)
		go c.HandlerTTSToClient(ctx, body, "", session, &wg)

	} else {
		talkRespCh := make(chan data.TalkResp, 10)
		if err := c.StreamingTalkByText(ctx, body, session, talkRespCh); err != nil {
			return fmt.Errorf("talk by text error: %s", err.Error())
		}

		for resp := range talkRespCh {

			if err := session.SendingMsgToClient(applet.ServiceType_Service_Nlp, resp, false, ""); err != nil {
				return err
			}
			for _, answer := range resp.AnsItem {
				ttsText := strings.Replace(strings.TrimSpace(answer.Text), "\n", "", -1)
				log.Debugf("sessionId:%s, ttsText:%s; start to call tts v2", session.TraceId, ttsText)
				if ttsText == "" {
					continue
				}
				wg.Add(1)
				go c.HandlerTTSToClient(ctx, ttsText, answer.Lang, session, &wg)
			}
		}
	}
	wg.Wait()
	if err := session.SendingMsgToClient(applet.ServiceType_Service_Nlp, "", true, ""); err != nil {
		return err
	}

	c.Debugf("HandlerText finished;")
	return nil
}

func (c *ChatService) HandlerTTSToClient(ctx context.Context, ttsText, language string, session *data.Session, wg *sync.WaitGroup) {
	defer func() {
		c.Infof("sessionId:%s,ttsText:%s, tts finished", session.TraceId, ttsText)
		wg.Done()
	}()
	ttsParam := session.TtsParam.Load().(*data.TTSParam)
	log.Debugf("start to call tts; sessionId:%s, ttsText:%s, ttsParam:%+v", session.TraceId, ttsText, ttsParam)
	ttsChan, err := c.CallTTSV2(ctx, session.Username, ttsParam, ttsText, language, session.Id, session.TraceId)
	if err != nil {
		c.Errorf("sessionId:%s,ttsText:%s, call tts error:%v", session.TraceId, ttsText, err)
		return
	}
	for ttsResp := range ttsChan {
		if err := session.SendingMsgToClient(applet.ServiceType_Service_TTS, base64.RawStdEncoding.EncodeToString(ttsResp), false, ""); err != nil {
			c.Errorf("sessionId:%s,ttsText:%s, send tts error:%v", session.TraceId, ttsText, err)
		}
	}
}
