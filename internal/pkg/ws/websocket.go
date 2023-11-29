package ws

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	mutex sync.RWMutex
)

func SendingMsgToClient(conn *websocket.Conn, msgType applet.ServiceType, content interface{}) error {
	log.Debugf("sending message to Client;msgType %s ", msgType)
	mutex.Lock()
	defer mutex.Unlock()
	return conn.WriteJSON(data.ChatServerMessage{
		ServiceType: msgType,
		Content:     content,
		IsEnd:       false,
		IsSuccess:   true,
		ErrMsg:      "",
	})
}

func SendFinishedMsgToClient(conn *websocket.Conn, msgType applet.ServiceType, content interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()
	log.Debugf("finished to send  message to client; message type:%s", msgType)
	return conn.WriteJSON(data.ChatServerMessage{
		ServiceType: msgType,
		Content:     content,
		IsEnd:       true,
		IsSuccess:   true,
		ErrMsg:      "",
	})
}

func SendErrMsgToClient(conn *websocket.Conn, msgType applet.ServiceType, errMsg string) error {
	mutex.Lock()
	defer mutex.Unlock()
	log.Debugf("finished to send  message to client; message type:%s", msgType)
	return conn.WriteJSON(data.ChatServerMessage{
		ServiceType: msgType,
		IsEnd:       true,
		IsSuccess:   false,
		ErrMsg:      errMsg,
	})
}
