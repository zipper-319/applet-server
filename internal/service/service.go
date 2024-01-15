package service

import (
	"applet-server/internal/data"
	"context"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewVoiceDataOperationService, NewAccountService, NewCloneSpeakerService, NewTTSServiceService, NewChatService, NewFeedbackService)

type ChatWebsocketServer interface {
	HandlerVoice(ctx context.Context, body chan []byte, session *data.Session) error
	HandlerText(ctx context.Context, body string, session *data.Session) error
}
