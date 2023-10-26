package server

import (
	"applet-server/internal/data"
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
)

var OperationChatWS = "/v2/ws/chat"

type SessionKey struct{}

type ChatWebsocketServer interface {
	HandlerVoice(ctx context.Context, body chan []byte, conn *websocket.Conn, session *data.Session) error
	HandlerText(ctx context.Context, body string, conn *websocket.Conn, session *data.Session) error
}

func RegisterChatWebsocketServer(s *http.Server, charService ChatWebsocketServer) {
	r := s.Route("/")
	r.GET(OperationChatWS, ChatWebsocketHandler(charService))
}

func ChatWebsocketHandler(srv ChatWebsocketServer) func(ctx http.Context) error {
	return nil
	//return func(ctx http.Context) error {
	//	var in applet.ChatWSReq
	//	if err := ctx.BindQuery(&in); err != nil {
	//		return err
	//	}
	//	if in.RobotId <= 0 {
	//		return errors.New("robotId is empty")
	//	}
	//	http.SetOperation(ctx, OperationChatWS)
	//	subProtocol := ctx.Header().Get("Sec-WebSocket-Protocol")
	//	log.Debugf("token:%s", subProtocol)
	//	h := ctx.Middleware(func(subCtx context.Context, req interface{}) (interface{}, error) {
	//
	//		upgrade := NewWsUpgrade(subProtocol)
	//		conn, err := upgrade.Upgrade(ctx.Response(), ctx.Request(), nil)
	//		if err != nil {
	//			fmt.Errorf("[websocket] fail to create ws, error: %v", err)
	//			return nil, err
	//		}
	//		log.Infof("[websocket] connect from %s", conn.RemoteAddr().String())
	//		defer conn.Close()
	//		isEndLastSession := make(chan struct{})
	//
	//		// 当建立一个连接时，就新起一个vad
	//		vadInputCh := make(chan []byte, 50)
	//		vadOutChan := make(chan []byte, 50)
	//
	//		traceId := uuid.New().String()
	//		session := data.GenSession(in, traceId)
	//		vadDateInfo := &vad.DataInfo{
	//			InputCh:  vadInputCh,
	//			OutputCh: vadOutChan,
	//			Conn:     conn,
	//			Session:  session,
	//			Server:   srv,
	//			IsEnd:    isEndLastSession,
	//		}
	//		ctxVad, cancelVad := context.WithCancel(context.Background())
	//		defer cancelVad()
	//		go vad.FeedAudio(ctxVad, vadDateInfo)
	//		awaitTime := 15 * time.Second
	//		vadTimer := time.NewTimer(awaitTime)
	//		var cancelText context.CancelFunc
	//		defer func() {
	//			vadTimer.Stop()
	//
	//			err := srv.StopDHuman(context.Background(), session)
	//			log.Debugf("[websocket] close conn; handle defer; sessionId: %s; err:%v", in.SessionId, err)
	//		}()
	//
	//		go func() {
	//			select {
	//			case <-vadTimer.C:
	//				// 触发超时机制，
	//				log.Infof("[websocket] vad timeout, close conn;sessionId: %s", in.SessionId)
	//				conn.Close()
	//			}
	//		}()
	//
	//		for {
	//
	//			messageType, message, err := conn.ReadMessage()
	//			if err != nil {
	//				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	//					log.Errorf("[websocket] read message error: %v", err)
	//				}
	//				log.Errorf("[websocket] read message error: %v", err)
	//				return nil, err
	//			}
	//			log.Debugf("sessionId:%s;reset:%t", in.SessionId, vadTimer.Reset(awaitTime))
	//
	//			switch messageType {
	//			case websocket.CloseMessage:
	//				return nil, nil
	//			case websocket.BinaryMessage:
	//				if err := conn.WriteMessage(websocket.BinaryMessage, []byte("")); err != nil {
	//					fmt.Errorf("[websocket] write pong message error: %v", err)
	//					return nil, err
	//				}
	//			case websocket.PingMessage:
	//				if err := conn.WriteMessage(websocket.PongMessage, []byte("")); err != nil {
	//					fmt.Errorf("[websocket] write pong message error: %v", err)
	//					return nil, err
	//				}
	//				log.Debug("[websocket] ping")
	//				break
	//			case websocket.PongMessage:
	//				log.Debug("[websocket] pong")
	//				break
	//			case websocket.TextMessage:
	//
	//				var chatMsg applet.ChatClientMessage
	//				//log.Debugf("message: %s", message)
	//				if err := json.Unmarshal(message, &chatMsg); err != nil {
	//					log.Errorf("sessionId:%s; err:%v", in.SessionId, err)
	//					if strings.HasPrefix(string(message), "text:") {
	//						text := strings.Trim(string(message), "text:")
	//						if err := srv.HandlerText(ctx, text, conn, data.GenSession(in, uuid.New().String())); err != nil {
	//							return nil, err
	//						}
	//					}
	//					break
	//				}
	//				if chatMsg.Type == applet.MessageType_chat_text {
	//					if vadDateInfo.LastCancel != nil {
	//						vadDateInfo.LastCancel()
	//						err := srv.CancelHandlerDHuman(context.Background(), session)
	//						log.Debugf("Cancel last HandlerDHuman; err=%v", err)
	//					}
	//					if cancelText != nil {
	//						cancelText()
	//						err := srv.CancelHandlerDHuman(context.Background(), session)
	//						log.Debugf(" Cancel last text HandlerDHuman; err=%v", err)
	//					}
	//					ctxText := context.Background()
	//					ctxText, cancelText = context.WithCancel(ctxText)
	//					session.TraceId = chatMsg.TraceId
	//					vadDateInfo.LastCancel = cancelText
	//					go srv.HandlerText(ctxText, chatMsg.Content, conn, session)
	//				}
	//				if chatMsg.Type == applet.MessageType_chat_voice {
	//
	//					if vadInputCh != nil {
	//						voice, err := base64.StdEncoding.DecodeString(chatMsg.Content)
	//						if err != nil {
	//							log.Errorf("[websocket] decode voice error: %v\n", err)
	//						} else {
	//							log.Debugf("[websocket] traceId:%s;voice length: %d\n", chatMsg.TraceId, len(voice))
	//							if len(voice) > 0 {
	//								vadInputCh <- voice
	//							}
	//						}
	//					}
	//				}
	//				if chatMsg.Type == applet.MessageType_chat_interruption {
	//					if cancelText != nil {
	//						cancelText()
	//					}
	//					if vadDateInfo.LastCancel != nil {
	//						vadDateInfo.LastCancel()
	//					}
	//					session.TraceId = chatMsg.TraceId
	//					err := srv.CancelHandlerDHuman(context.Background(), session)
	//					log.Debugf("[websocket] interruption;sessionId:%s, traceId:%s, err:%v", in.SessionId, chatMsg.TraceId, err)
	//					if err != nil {
	//						ws.SendFinishedMsgToClient(conn, applet.ServiceType_service_digital_human, "success to interrupt")
	//					}
	//				}
	//
	//				if chatMsg.Type == applet.MessageType_chat_parameter {
	//					var parameter applet.ChatParameter
	//					if err := json.Unmarshal([]byte(chatMsg.Content), &parameter); err != nil {
	//						log.Errorf("[websocket] unmarshal parameter error: %v\n", err)
	//						ws.SendErrMsgToClient(conn, applet.ServiceType_Service_VAD, err.Error())
	//					} else {
	//						session.Language.Store(parameter.Language)
	//					}
	//				}
	//
	//				break
	//			}
	//
	//		}
	//	})
	//	out, err := h(ctx, &in)
	//	if err != nil {
	//		return err
	//	}
	//	reply := out.(*applet.ChatWSResp)
	//	return ctx.Result(200, reply)
	//}
}

func NewWsUpgrade(subProtocol string) websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{subProtocol},
	}
}
