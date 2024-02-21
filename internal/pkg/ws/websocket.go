package ws

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/pkg/log"
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

type WsClient struct {
	*websocket.Conn
	sync.RWMutex
	*log.MyLogger
}

type ChatServerMessage struct {
	ServiceType applet.ServiceType `json:"service_type"`
	QuestionId  string             ` json:"question_id"`
	SessionId   string             ` json:"session_id"`
	Content     interface{}        ` json:"content"`
	IsEnd       bool               `json:"is_end"`
	IsSuccess   bool               `json:"is_success"`
	ErrMsg      string             `json:"err_msg"`
}

func NewWsClient(conn *websocket.Conn, logger *log.MyLogger) *WsClient {
	return &WsClient{
		Conn:     conn,
		MyLogger: logger,
	}
}

func (c *WsClient) SendingMsgToClient(ctx context.Context, msgType applet.ServiceType, content interface{}, isEnd bool, errMsg string) error {

	c.WithContext(ctx).Debugf("sending message to Client;msgType %s ", msgType)
	questionId, _ := ctx.Value("questionId").(string)
	sessionId, _ := ctx.Value("sessionId").(string)
	c.Lock()
	defer c.Unlock()
	if errMsg != "" {
		return c.WriteJSON(ChatServerMessage{
			ServiceType: msgType,
			QuestionId:  questionId,
			SessionId:   sessionId,
			Content:     nil,
			IsEnd:       true,
			IsSuccess:   false,
			ErrMsg:      errMsg,
		})
	}
	return c.WriteJSON(ChatServerMessage{
		ServiceType: msgType,
		QuestionId:  questionId,
		SessionId:   sessionId,
		Content:     content,
		IsEnd:       isEnd,
		IsSuccess:   true,
		ErrMsg:      "",
	})
}
