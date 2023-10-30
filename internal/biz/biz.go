package biz

import (
	"applet-server/internal/biz/tts"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewS3UseCase, NewUserUseCase, tts.NewTTSClient, NewCloneSpeakerUseCase)
