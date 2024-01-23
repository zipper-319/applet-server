package service

import (
	"applet-server/internal/biz"
	"applet-server/internal/biz/tts"
	jwtUtil "applet-server/internal/pkg/jwt"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "applet-server/api/v2/applet"
)

type TTSServiceService struct {
	pb.UnimplementedTTSServiceServer
	ttsClient *tts.TTSClient
	speakerUC *biz.CloneSpeakerUseCase
}

func NewTTSServiceService(client *tts.TTSClient, useCase *biz.CloneSpeakerUseCase) *TTSServiceService {
	return &TTSServiceService{
		ttsClient: client,
		speakerUC: useCase,
	}
}

func (s *TTSServiceService) GetTTSConfig(ctx context.Context, req *emptypb.Empty) (*pb.GetTTSConfigResult, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	ttsConfig, err := s.ttsClient.GetTTSConfig(ctx, tokenInfo.Username)
	if err != nil {
		return nil, err
	}
	pitchList := make([]*pb.MessagePitch, 0, len(ttsConfig.PitchList))
	for _, pitch := range ttsConfig.PitchList {
		pitchList = append(pitchList, &pb.MessagePitch{
			Name:        pitch.Name,
			ChineseName: pitch.ChineseName,
		})
	}
	emotionList := make([]*pb.MessageEmotion, 0, len(ttsConfig.EmotionList))
	for _, emotion := range ttsConfig.EmotionList {
		emotionList = append(emotionList, &pb.MessageEmotion{
			Name:        emotion.Name,
			ChineseName: emotion.ChineseName,
		})
	}
	movementList := make([]*pb.MessageMovement, 0, len(ttsConfig.MovementList))
	for _, movement := range ttsConfig.MovementList {
		movementList = append(movementList, &pb.MessageMovement{
			Name:        movement.Name,
			ChineseName: movement.ChineseName,
		})
	}
	expressionList := make([]*pb.MessageExpression, 0, len(ttsConfig.ExpressionList))
	for _, expression := range ttsConfig.ExpressionList {
		expressionList = append(expressionList, &pb.MessageExpression{
			Name:        expression.Name,
			ChineseName: expression.ChineseName,
		})
	}
	cloneSpeakerMap, err := s.speakerUC.GetSpeakerMap(ctx, tokenInfo.Username)
	if err != nil {
		return nil, err
	}

	speakerList := make([]*pb.SpeakerParameter, 0, len(ttsConfig.SpeakerList.List))
	for _, speaker := range ttsConfig.SpeakerList.List {
		temp := &pb.SpeakerParameter{
			SpeakerName:          speaker.SpeakerName,
			ParameterSpeakerName: speaker.ParameterSpeakerName,
			IsSupportEmotion:     speaker.IsSupportEmotion,
			IsSupportMixedVoice:  speaker.IsSupportMixedVoice,
			IsBelongClone:        speaker.IsBelongClone,
		}
		if speaker.IsBelongClone {
			if speakerName, exist := cloneSpeakerMap[temp.ParameterSpeakerName]; exist {
				temp.ParameterSpeakerName = speakerName
			} else {
				continue
			}
		}
		speakerList = append(speakerList, temp)
	}

	return &pb.GetTTSConfigResult{
		SpeakerList:    speakerList,
		SpeedList:      ttsConfig.SpeedList,
		VolumeList:     ttsConfig.VolumeList,
		PitchList:      pitchList,
		EmotionList:    emotionList,
		MovementList:   movementList,
		ExpressionList: expressionList,
	}, nil
}
