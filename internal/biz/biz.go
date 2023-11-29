package biz

import (
	"applet-server/internal/biz/asr"
	"applet-server/internal/biz/nlp"
	"applet-server/internal/biz/tts"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewVideoUseCase, NewUserUseCase, tts.NewTTSClient, NewCloneSpeakerUseCase, asr.NewAsRControllerClient, nlp.NewTalkClient)
