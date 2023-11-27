package biz

import (
	"applet-server/internal/biz/tts"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewVideoUseCase, NewUserUseCase, tts.NewTTSClient, NewCloneSpeakerUseCase)
