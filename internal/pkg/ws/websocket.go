package ws

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/pkg/log"
	"github.com/gorilla/websocket"
	"sync"
)

type WsClient struct {
	*websocket.Conn
	sync.RWMutex
}

type ChatServerMessage struct {
	ServiceType applet.ServiceType ` json:"service_type,omitempty"`
	Content     interface{}        ` json:"content,omitempty"`
	IsEnd       bool               `son:"is_end,omitempty"`
	IsSuccess   bool               `json:"is_success,omitempty"`
	ErrMsg      string             `json:"err_msg,omitempty"`
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{Conn: conn}
}

func (c *WsClient) SendingMsgToClient(msgType applet.ServiceType, content interface{}, isEnd bool, errMsg string) error {
	log.Debugf("sending message to Client;msgType %s ", msgType)
	c.Lock()
	defer c.Unlock()
	if errMsg != "" {
		return c.WriteJSON(ChatServerMessage{
			ServiceType: msgType,
			IsEnd:       true,
			IsSuccess:   false,
			ErrMsg:      errMsg,
		})
	}
	return c.WriteJSON(ChatServerMessage{
		ServiceType: msgType,
		Content:     content,
		IsEnd:       isEnd,
		IsSuccess:   true,
		ErrMsg:      "",
	})
}
