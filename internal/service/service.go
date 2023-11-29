package service

import (
	"applet-server/internal/data"
	"context"
	"github.com/google/wire"
	"github.com/gorilla/websocket"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewVoiceDataOperationService, NewAccountService, NewCloneSpeakerService, NewTTSServiceService, NewChatService)

type ChatWebsocketServer interface {
	HandlerVoice(ctx context.Context, body chan []byte, conn *websocket.Conn, session *data.Session) error
	HandlerText(ctx context.Context, body string, conn *websocket.Conn, session *data.Session) error
}
