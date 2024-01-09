package server

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/ws"
	"applet-server/internal/vad"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

var OperationChatWS = "/api/v2/ws/chat"

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

	return func(ctx http.Context) error {
		var in applet.ChatWSReq
		log.Debugf("ChatWebsocketHandler: %v", ctx.Request().Header)
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		log.Debugf("ChatWebsocketHandler: %v", in)
		if in.Method == applet.MethodType_TypeUnknown {
			return errors.New("method type is empty")
		}
		http.SetOperation(ctx, OperationChatWS)
		subProtocol := ctx.Header().Get("Sec-WebSocket-Protocol")
		token := subProtocol
		log.Debugf("token:%s", token)
		tokenInfo, err := jwtUtil.ParseToken(token, "")
		if err != nil {
			return err
		}

		h := ctx.Middleware(func(subCtx context.Context, req interface{}) (interface{}, error) {

			upgrade := NewWsUpgrade(subProtocol)
			conn, err := upgrade.Upgrade(ctx.Response(), ctx.Request(), nil)
			if err != nil {
				fmt.Errorf("[websocket] fail to create ws, error: %v", err)
				return nil, err
			}
			log.Infof("[websocket] connect from %s", conn.RemoteAddr().String())
			defer conn.Close()
			sessionId := uuid.New().String()
			session := data.GenSession(in, tokenInfo.Username, sessionId)
			ttsParam := &data.TTSParam{
				Speaker: "DaXiaoFang",
				Speed:   "3",
				Volume:  "3",
			}
			session.TtsParam.Store(ttsParam)

			awaitTime := 15 * time.Second
			connectTimer := time.NewTimer(awaitTime)

			var vadInputCh chan []byte
			var cancelText context.CancelFunc
			var vadDateInfo *vad.DataInfo

			// 当建立一个连接时，就新起一个vad
			if in.Method >= applet.MethodType_OnlyAsr {
				vadInputCh = make(chan []byte, 50)
				vadDateInfo = &vad.DataInfo{
					InputCh: vadInputCh,
					Conn:    conn,
					Session: session,
					Server:  srv,
				}
				ctxVad, cancelVad := context.WithCancel(context.Background())
				defer cancelVad()
				go vad.FeedAudio(ctxVad, vadDateInfo)

				defer func() {
					connectTimer.Stop()
					log.Debugf("[websocket] close conn; handle defer; sessionId: %s; err:%v", in.SessionId, err)
				}()
			}

			//go func() {
			//	select {
			//	case <-connectTimer.C:
			//		// 触发超时机制，
			//		log.Infof("[websocket] vad timeout, close conn;sessionId: %s", in.SessionId)
			//		conn.Close()
			//	}
			//}()

			for {

				messageType, message, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Errorf("[websocket] read message error: %v", err)
					}
					log.Errorf("[websocket] read message error: %v", err)
					return nil, err
				}
				log.Debugf("sessionId:%s;reset:%t", session.Id, connectTimer.Reset(awaitTime))

				switch messageType {
				case websocket.CloseMessage:
					return nil, nil
				case websocket.BinaryMessage:
					if vadInputCh != nil && len(message) > 0 {
						vadInputCh <- message
					}
				case websocket.PingMessage:
					if err := conn.WriteMessage(websocket.PongMessage, []byte("")); err != nil {
						fmt.Errorf("[websocket] write pong message error: %v", err)
						return nil, err
					}
					log.Debug("[websocket] ping")
					break
				case websocket.PongMessage:
					log.Debug("[websocket] pong")
					break
				case websocket.TextMessage:

					var chatMsg applet.ChatClientMessage
					//log.Debugf("message: %s", message)
					// use to debug
					if err := json.Unmarshal(message, &chatMsg); err != nil {
						log.Errorf("message:%s, sessionId:%s; err:%v", message, in.SessionId, err)
						if strings.HasPrefix(string(message), "text:") {
							text := strings.Trim(string(message), "text:")
							session.TraceId = chatMsg.TraceId
							if err := srv.HandlerText(ctx, text, conn, session); err != nil {
								return nil, err
							}
						}
						break
					}
					log.Debugf("chatMsg:%v", chatMsg)
					if chatMsg.MessageType == applet.MessageType_chat_text {

						ctxText := context.Background()
						go srv.HandlerText(ctxText, chatMsg.Content, conn, session)
					}
					if chatMsg.MessageType == applet.MessageType_chat_voice {

						if vadInputCh != nil {
							voice, err := base64.StdEncoding.DecodeString(chatMsg.Content)
							if err != nil {
								log.Errorf("[websocket] decode voice error: %v\n", err)
							} else {
								log.Debugf("[websocket] traceId:%s;voice length: %d\n", chatMsg.TraceId, len(voice))
								if len(voice) > 0 {
									vadInputCh <- voice
								}
							}
						}
					}
					if chatMsg.MessageType == applet.MessageType_chat_interruption {
						if cancelText != nil {
							cancelText()
						}
						if vadDateInfo.LastCancel != nil {
							vadDateInfo.LastCancel()
						}
						session.TraceId = chatMsg.TraceId
						log.Debugf("[websocket] interruption;sessionId:%s, traceId:%s, err:%v", in.SessionId, chatMsg.TraceId, err)
					}

					if chatMsg.MessageType == applet.MessageType_chat_parameter {
						var parameter applet.ChatParameter
						if err := json.Unmarshal([]byte(chatMsg.Content), &parameter); err != nil {
							log.Errorf("[websocket] unmarshal parameter error: %v\n", err)
							ws.SendErrMsgToClient(conn, applet.ServiceType_Service_VAD, err.Error())
						} else {

							log.Debugf("[websocket] parameter;sessionId:%s, traceId:%s, parameter:%v", in.SessionId, chatMsg.TraceId, parameter)
							ttsParamNew := session.TtsParam.Load().(*data.TTSParam)
							if parameter.Pitch != "" {
								ttsParamNew.Pitch = parameter.Pitch
							}
							if parameter.Speed != "" {
								ttsParamNew.Speed = parameter.Speed
							}
							if parameter.Volume != "" {
								ttsParamNew.Volume = parameter.Volume
							}
							if parameter.Speaker != "" {
								ttsParamNew.Speaker = parameter.Speaker
							}
							if parameter.IsClone > 0 {
								if parameter.IsClone == 1 {
									ttsParamNew.IsClone = false
								} else {
									ttsParamNew.IsClone = true
								}
							}
							session.TtsParam.Store(ttsParamNew)
						}
					}

					break
				}

			}
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*applet.ChatWSResp)
		return ctx.Result(200, reply)
	}
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
