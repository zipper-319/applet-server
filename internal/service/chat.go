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
	"runtime/debug"
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
			c.WithContext(ctx).Errorf("HandlerVoice; err:%v", err)
			fmt.Println(string(debug.Stack()))
		}
		cancel()
		c.WithContext(ctx).Debug("-------------end to  handler Voice----------------")
	}()
	c.WithContext(ctx).Debug("--------------start to  handler Voice-----------")

	asrOutChan := make(chan string, 10)
	go c.AsRControllerClient.StreamingRecognize(ctx, session, vadOutChan, asrOutChan)

	asrResult := ""
	c.WithContext(ctx).Debug("await to asrResult ")
	for asrText := range asrOutChan {
		asrResult = asrText
		c.WithContext(ctx).Debugf("asrResult:%s", asrResult)
		if asrResult != "" {
			session.SendingMsgToClient(ctx, applet.ServiceType_Service_ASR, asrText, false, "")
		}
	}
	session.SendingMsgToClient(ctx, applet.ServiceType_Service_ASR, asrResult, true, "")
	if asrResult == "" {
		return fmt.Errorf("asrResult is empty")
	}

	talkRespCh := make(chan data.TalkResp, 10)
	if err := c.TalkClient.StreamingTalkByText(ctx, asrResult, session, talkRespCh); err != nil {
		return fmt.Errorf("talk by text error: %s", err.Error())
	}

	var wg sync.WaitGroup

	for resp := range talkRespCh {
		c.WithContext(ctx).Debugf("resp:%v", resp)
		if err := session.SendingMsgToClient(ctx, applet.ServiceType_Service_Nlp, resp, false, ""); err != nil {
			return err
		}
		for _, answer := range resp.AnsItem {
			ttsText := strings.Replace(strings.TrimSpace(answer.Text), "\n", "", -1)
			c.WithContext(ctx).Debugf("ttsText:%s; start to call tts v2", ttsText)
			if ttsText == "" {
				continue
			}
			wg.Add(1)
			go c.HandlerTTSToClient(ctx, ttsText, answer.Lang, session, &wg)

		}
	}
	wg.Wait()
	if err := session.SendingMsgToClient(ctx, applet.ServiceType_Service_Nlp, "", true, ""); err != nil {
		return err
	}
	c.WithContext(ctx).Debugf("the sentence finished")
	return nil
}
func (c *ChatService) HandlerText(ctx context.Context, body string, session *data.Session) error {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	c.WithContext(ctx).Infof("HandlerText;text:%s; env:%s, username:%s", body, session.Env, session.Username)
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

			if err := session.SendingMsgToClient(ctx, applet.ServiceType_Service_Nlp, resp, false, ""); err != nil {
				return err
			}
			for _, answer := range resp.AnsItem {
				ttsText := strings.Replace(strings.TrimSpace(answer.Text), "\n", "", -1)
				c.WithContext(ctx).Debugf("ttsText:%s; start to call tts v2", ttsText)
				if ttsText == "" {
					continue
				}
				wg.Add(1)
				go c.HandlerTTSToClient(ctx, ttsText, answer.Lang, session, &wg)
			}
		}
	}
	wg.Wait()
	if err := session.SendingMsgToClient(ctx, applet.ServiceType_Service_Nlp, "", true, ""); err != nil {
		return err
	}

	c.WithContext(ctx).Debugf("HandlerText finished;")
	return nil
}

func (c *ChatService) HandlerTTSToClient(ctx context.Context, ttsText, language string, session *data.Session, wg *sync.WaitGroup) {
	defer func() {
		c.WithContext(ctx).Infof("ttsText:%s, tts finished", ttsText)
		wg.Done()
	}()
	ttsParam := session.TtsParam.Load().(*data.TTSParam)
	log.Debugf("start to call tts; ttsText:%s, ttsParam:%+v", ttsText, ttsParam)
	ttsChan, err := c.CallTTSV2(ctx, session, ttsParam, ttsText, language)
	if err != nil {
		c.WithContext(ctx).Errorf("ttsText:%s, call tts error:%v", ttsText, err)
		return
	}
	for ttsResp := range ttsChan {
		if err := session.SendingMsgToClient(ctx, applet.ServiceType_Service_TTS, base64.RawStdEncoding.EncodeToString(ttsResp), false, ""); err != nil {
			c.Errorf("ttsText:%s, send tts error:%v", ttsText, err)
		}
	}
	session.SendingMsgToClient(ctx, applet.ServiceType_Service_TTS, "", true, "")
}
