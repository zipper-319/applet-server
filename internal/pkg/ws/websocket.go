package ws

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	v1 "sv_capacity_control/api/chat/v1"
	"sv_capacity_control/internal/data"
	"sync"
)

// NewWebsocketServer create a websocket server.
//func NewWebsocketServer(c *conf.Server, _ log.Logger, svc *service.ChatRoomService) *websocket.Server {
//	srv := websocket.NewServer(
//		websocket.WithAddress(c.Websocket.Addr),
//		websocket.WithPath(c.Websocket.Path),
//		websocket.WithConnectHandle(svc.OnWebsocketConnect),
//		websocket.WithCodec("json"),
//	)
//
//	svc.SetWebsocketServer(srv)
//
//	srv.RegisterMessageHandler(websocket.MessageType(v1.MessageType_voice),
//		func(sessionId websocket.SessionID, payload websocket.MessagePayload) error {
//			switch t := payload.(type) {
//			case *v1.ChatMessage:
//				return svc.OnChatMessage(sessionId, t)
//			default:
//				return errors.New("invalid payload type")
//			}
//		},
//		func() websocket.Any { return &v1.ChatMessage{} },
//	)
//
//	return srv
//}


var(
	mutex sync.RWMutex
)


func SendingMsgToClient(conn *websocket.Conn, msgType v1.ServiceType, content interface{}) error {
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

func SendFinishedMsgToClient(conn *websocket.Conn, msgType v1.ServiceType, content interface{}) error {
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

func SendErrMsgToClient(conn *websocket.Conn, msgType v1.ServiceType, errMsg string) error {
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
