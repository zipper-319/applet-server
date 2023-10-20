package service

import (
	pb "applet-server/api/v2/applet"
	"applet-server/internal/biz"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/voiceText"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type VoiceDataOperationService struct {
	pb.UnimplementedVoiceDataOperationServer
	*biz.S3UseCase
	*log.Helper
}

func NewVoiceDataOperationService(useCase *biz.S3UseCase, logger log.Logger) *VoiceDataOperationService {
	return &VoiceDataOperationService{
		S3UseCase: useCase,
		Helper:    log.NewHelper(logger),
	}
}

func (s *VoiceDataOperationService) PutVoiceData(ctx context.Context, req *pb.VoiceDataReqData) (*pb.VoiceDataResData, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	s.Infof("tokenInfo: %+v", tokenInfo)
	username := tokenInfo.Username
	voiceData, err := base64.StdEncoding.DecodeString(req.Voice)
	if err != nil {
		return nil, err
	}
	if err := s.S3UseCase.Upload(ctx, voiceData, int(req.Sequence), username, int(req.SpeakerSerial), req.VoiceType); err != nil {
		return nil, err
	}
	return &pb.VoiceDataResData{}, nil
}
func (s *VoiceDataOperationService) GetProgress(ctx context.Context, req *pb.ProgressRequest) (*pb.ProgressResData, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	s.Infof("tokenInfo: %+v", tokenInfo)
	username := tokenInfo.Username
	key := fmt.Sprintf("finishedTime:%s:%s:%d", username, req.VoiceType, req.SpeakerSerial)
	finishedTime, err := s.S3UseCase.Data.RedisClient.Get(ctx, key).Int64()
	if err != nil {
		return nil, err
	}
	return &pb.ProgressResData{
		CurrentNumber: 0,
		FinishedTime:  finishedTime,
		AwaitTrain:    0,
	}, nil
}
func (s *VoiceDataOperationService) DownloadVoice(ctx context.Context, req *pb.DownloadReqData) (*pb.DownloadResData, error) {
	return &pb.DownloadResData{}, nil
}
func (s *VoiceDataOperationService) Commit(ctx context.Context, req *pb.CommitRequest) (*pb.CommitResData, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	s.Infof("tokenInfo: %+v", tokenInfo)
	username := tokenInfo.Username
	s3NameList := make([]string, 0, 50)
	for i := 0; i < voiceText.VoiceDataSize[req.VoiceType]; i++ {
		s3NameList = append(s3NameList, fmt.Sprintf("%s/%d/%d.pcm", username, req.SpeakerSerial, i))
	}
	key := fmt.Sprintf("finishedTime:%s:%s:%d", username, req.VoiceType, req.SpeakerSerial)
	s.S3UseCase.Data.RedisClient.Set(ctx, key, time.Now().Unix(), 0)
	return &pb.CommitResData{}, nil
}

func (s *VoiceDataOperationService) GetText(ctx context.Context, req *pb.GetTextRequest) (*pb.GetTextResData, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	s.Infof("tokenInfo: %+v", tokenInfo)
	content, err := voiceText.ReadText(req.VoiceType)
	if err != nil {
		return nil, err
	}
	return &pb.GetTextResData{
		Text: content,
	}, nil
}
