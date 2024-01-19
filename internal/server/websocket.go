package server

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/log"
	"applet-server/internal/vad"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

var OperationChatWS = "/api/v2/ws/chat"

type SessionKey struct{}

type ChatWebsocketServer interface {
	HandlerVoice(ctx context.Context, body chan []byte, session *data.Session) error
	HandlerText(ctx context.Context, body string, session *data.Session) error
}

func RegisterChatWebsocketServer(s *http.Server, charService ChatWebsocketServer, logger *log.MyLogger) {
	r := s.Route("/")
	r.GET(OperationChatWS, ChatWebsocketHandler(charService, logger))
}

func ChatWebsocketHandler(srv ChatWebsocketServer, logger *log.MyLogger) func(ctx http.Context) error {

	return func(ctx http.Context) error {
		var in applet.ChatWSReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		log.Debugf("ChatWSReq: %v", in)
		if in.Method == applet.MethodType_TypeUnknown {
			return errors.New("method type is empty")
		}
		http.SetOperation(ctx, OperationChatWS)
		subProtocol := ctx.Header().Get("Sec-WebSocket-Protocol")
		token := subProtocol
		logger.WithContext(ctx).Debugf("token:%s", token)
		tokenInfo, err := jwtUtil.ParseToken(token, "")
		if err != nil {
			return err
		}

		h := ctx.Middleware(func(subCtx context.Context, req interface{}) (interface{}, error) {

			upgrade := NewWsUpgrade(subProtocol)
			conn, err := upgrade.Upgrade(ctx.Response(), ctx.Request(), nil)
			if err != nil {
				logger.WithContext(subCtx).Errorf("[websocket] fail to create ws, error: %v", err)
				return nil, err
			}
			logger.WithContext(subCtx).Infof("[websocket] connect from %s", conn.RemoteAddr().String())
			defer conn.Close()
			sessionId := uuid.New().String()
			session := data.GenSession(in, tokenInfo.Username, sessionId, conn, logger)
			subCtx = context.WithValue(subCtx, "sessionId", sessionId)
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
					Session: session,
					Server:  srv,
				}
				ctxVad, cancelVad := context.WithCancel(context.Background())
				defer cancelVad()
				go vad.FeedAudio(ctxVad, vadDateInfo)

				defer func() {
					connectTimer.Stop()
					logger.WithContext(subCtx).Debugf("[websocket] close conn; handle defer; sessionId: %s; err:%v", err)
				}()
			}

			for {

				messageType, message, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						logger.WithContext(subCtx).Errorf("[websocket] read message error: %v", err)
					}
					logger.WithContext(subCtx).Errorf("[websocket] read message error: %v", err)
					return nil, err
				}
				logger.WithContext(subCtx).Debugf("sessionId:%s;reset:%t", session.Id, connectTimer.Reset(awaitTime))

				switch messageType {
				case websocket.CloseMessage:
					return nil, nil
				case websocket.BinaryMessage:
					if vadInputCh != nil && len(message) > 0 {
						vadInputCh <- message
					}
				case websocket.PingMessage:
					if err := conn.WriteMessage(websocket.PongMessage, []byte("")); err != nil {
						logger.WithContext(subCtx).Errorf("[websocket] write pong message error: %v", err)
						return nil, err
					}
					logger.WithContext(subCtx).Debug("[websocket] ping")
					break
				case websocket.PongMessage:
					logger.WithContext(subCtx).Debug("[websocket] pong")
					break
				case websocket.TextMessage:

					var chatMsg applet.ChatClientMessage
					if err := json.Unmarshal(message, &chatMsg); err != nil {
						logger.WithContext(subCtx).Errorf("message:%s, sessionId:%s; err:%v", message, sessionId, err)
						if strings.HasPrefix(string(message), "text:") {
							text := strings.Trim(string(message), "text:")
							questionId := uuid.New().String()
							ctx := context.WithValue(ctx, "questionId", questionId)
							if err := srv.HandlerText(ctx, text, session); err != nil {
								return nil, err
							}
						}
						break
					}
					logger.WithContext(subCtx).Debugf("chatMsg:%v", chatMsg)
					if chatMsg.MessageType == applet.MessageType_chat_text {

						go srv.HandlerText(subCtx, chatMsg.Content, session)
					}
					if chatMsg.MessageType == applet.MessageType_chat_voice {

						if vadInputCh != nil {
							voice, err := base64.StdEncoding.DecodeString(chatMsg.Content)
							if err != nil {
								logger.WithContext(subCtx).Errorf("[websocket] decode voice error: %v\n", err)
							} else {
								logger.WithContext(subCtx).Debugf("[websocket] voice length: %d\n", len(voice))
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
						logger.WithContext(subCtx).Debugf("[websocket] interruption; err:%v", err)
					}

					if chatMsg.MessageType == applet.MessageType_chat_parameter {
						var parameter applet.ChatParameter
						if err := json.Unmarshal([]byte(chatMsg.Content), &parameter); err != nil {
							logger.WithContext(subCtx).Errorf("[websocket] unmarshal parameter error: %v\n", err)
							session.SendingMsgToClient(ctx, applet.ServiceType_Service_VAD, "", true, err.Error())
						} else {

							logger.WithContext(subCtx).Debugf("[websocket] parameter; parameter:%v", parameter)
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
