package nlp

import (
	"applet-server/internal/biz/nlp/proto"
	"applet-server/internal/conf"
	"applet-server/internal/data"
	"applet-server/internal/pkg/log"
	"applet-server/internal/pkg/util"
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/idoubi/goz"
	"google.golang.org/grpc"
	"strconv"
	"strings"
)

type TalkClient struct {
	*grpc.ClientConn
	*log.MyLogger
}

func NewTalkClient(c *conf.App, logger *log.MyLogger) *TalkClient {
	ctx, _ := context.WithTimeout(context.Background(), c.Asr.GetTimeout().AsDuration())
	conn, err := grpc.DialContext(ctx, c.Nlp.GetAddr(),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	return &TalkClient{
		ClientConn: conn,
		MyLogger:   logger,
	}
}

func (c *TalkClient) StreamingTalk(ctx context.Context, session *data.Session, questionStream chan string) (talkRespCh chan data.TalkResp, err error) {
	talkRespCh = make(chan data.TalkResp, 10)
	streamingTalkClient := pb.NewTalkClient(c.ClientConn)
	streamingTalk, err := streamingTalkClient.StreamingTalk(ctx)

	if err != nil {
		return
	}
	go TalkReceive(ctx, streamingTalk, talkRespCh, session.TraceId)
	question := ""

	for q := range questionStream {
		question = q
		req := genTalkRequest(false, question, session)
		err = streamingTalk.Send(req)
		if err != nil {
			return
		}
	}
	req := genTalkRequest(true, question, session)
	err = streamingTalk.Send(req)
	if err != nil {
		return
	}
	streamingTalk.CloseSend()
	return
}

func (c *TalkClient) StreamingTalkByText(ctx context.Context, question string, session *data.Session, talkRespCh chan data.TalkResp) error {

	streamingTalkClient := pb.NewTalkClient(c.ClientConn)
	streamingTalk, err := streamingTalkClient.StreamingTalk(ctx)
	if err != nil {
		return err
	}
	req := genTalkRequest(true, question, session)
	if err = streamingTalk.Send(req); err != nil {
		return err
	}
	if err := streamingTalk.CloseSend(); err != nil {
		return err
	}
	go TalkReceive(ctx, streamingTalk, talkRespCh, session.TraceId)
	log.Debugf("sessionId:%s, text:%s; finish to call streamingTalk", session.TraceId, question)
	return nil
}

func genTalkRequest(isFull bool, question string, session *data.Session) *pb.TalkRequest {
	questionId := session.TraceId
	sessionId := session.Id
	agentId := int64(session.AgentId)
	robotId := strconv.Itoa(int(session.RobotId))
	position := session.Position
	language := session.Language.Load()

	return &pb.TalkRequest{
		IsFull:     isFull,
		AgentID:    agentId,
		SessionID:  sessionId,
		QuestionID: questionId,
		EventType:  pb.Text,
		EnvInfo:    make(map[string]string),
		RobotID:    robotId,
		TenantCode: "mpTenantId",
		Version:    "v3", //speech.Header.Version,
		Position:   position,
		Asr: pb.Asr{
			Lang: language,
			Text: question,
		},
	}
}

func TalkReceive(ctx context.Context, streamingTalk pb.Talk_StreamingTalkClient, talkRespCh chan data.TalkResp, sessionId string) {
	defer func() {
		log.Debugf("finish to receive  streamTalk")
		close(talkRespCh)
	}()
	for {
		select {
		case <-ctx.Done():
			log.Debugf("TalkReceive: %s had canceled", sessionId)
			return
		default:
			resp, err := streamingTalk.Recv()
			if err != nil {
				return
			}
			log.Debugf("sessionId:%s;TalkReceive receive response: %s, cost:%d", sessionId, resp.Source, resp.Cost)
			talkRespCh <- parseTalkResp(resp)
		}
	}
}

func parseTalkResp(resp *pb.TalkResponse) data.TalkResp {
	talkResp := data.TalkResp{}
	if resp != nil {
		talkResp.Source = resp.Source
		if len(resp.Tts) > 0 {
			talkResp.AnsItem = make([]*data.Answer, 0)

			for i := range resp.Tts {
				var ansItem *data.Answer
				if resp.Tts[i].Text != "" {
					if ansItem == nil {
						ansItem = new(data.Answer)
					}
					ansItem.Text = strings.ReplaceAll(strings.TrimSpace(resp.Tts[i].Text), "\n", "")
					ansItem.Lang = resp.Tts[i].Lang
				}
				if resp.Tts[i].Action.Param.Url != "" {
					if ansItem == nil {
						ansItem = new(data.Answer)
					}
					ansItem.MusicUrl = resp.Tts[i].Action.Param.Url
				}
				if resp.Tts[i].Action.Param.VideoUrl != "" {
					if ansItem == nil {
						ansItem = new(data.Answer)
					}
					ansItem.VideoUrl = resp.Tts[i].Action.Param.VideoUrl
				}
				if resp.Tts[i].Action.Param.PicUrl != "" {
					if ansItem == nil {
						ansItem = new(data.Answer)
					}
					ansItem.PicUrl = resp.Tts[i].Action.Param.PicUrl
				}
				if resp.Tts[i].Payload != "" {

					if payload, err := simplejson.NewJson([]byte(resp.Tts[i].Payload)); err == nil {
						mediaApi := payload.Get("media_api")
						url := fmt.Sprintf("%s%s", util.MediaApiUrl, mediaApi)
						if resp, err := goz.Get(url); err == nil {
							if body, err := resp.GetBody(); err == nil {
								if mediaApiResult, err := simplejson.NewJson(body); err == nil {
									if videoUrl, err := mediaApiResult.Get("data").Get("videourl").String(); err == nil && videoUrl != "" {
										if ansItem == nil {
											ansItem = new(data.Answer)
										}
										ansItem.VideoUrl = videoUrl
									}
									if picUrl, err := mediaApiResult.Get("data").Get("img_url").String(); err == nil && picUrl != "" {
										if ansItem == nil {
											ansItem = new(data.Answer)
										}
										ansItem.VideoUrl = picUrl
									}
								}
							}
						}
					}

				}

				if ansItem != nil {
					talkResp.AnsItem = append(talkResp.AnsItem, ansItem)
				}
			}
		}
		if resp.HitLog != nil {
			talkResp.Domain = resp.HitLog.Fields["domain"].GetStringValue()
			talkResp.Intent = resp.HitLog.Fields["intent"].GetStringValue()
		}

	}
	return talkResp
}
