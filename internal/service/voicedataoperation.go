package service

import (
	pb "applet-server/api/v2/applet"
	"applet-server/internal/biz"
	"applet-server/internal/pkg/voiceText"
	"context"
	"encoding/base64"
	"fmt"
	"time"
)

type VoiceDataOperationService struct {
	pb.UnimplementedVoiceDataOperationServer
	*biz.S3UseCase
}

func NewVoiceDataOperationService(useCase *biz.S3UseCase) *VoiceDataOperationService {
	return &VoiceDataOperationService{
		S3UseCase: useCase,
	}
}

func (s *VoiceDataOperationService) PutVoiceData(ctx context.Context, req *pb.VoiceDataReqData) (*pb.VoiceDataResData, error) {
	var username string
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
	var username string
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
	var username string
	s3NameList := make([]string, 0, 50)
	for i := 0; i < voiceText.VoiceDataSize[req.VoiceType.String()]; i++ {
		s3NameList = append(s3NameList, fmt.Sprintf("%s/%d/%d.pcm", username, req.SpeakerSerial, i))
	}
	key := fmt.Sprintf("finishedTime:%s:%s:%d", username, req.VoiceType, req.SpeakerSerial)
	s.S3UseCase.Data.RedisClient.Set(ctx, key, time.Now().Unix(), 0)
	return &pb.CommitResData{}, nil
}

func (s *VoiceDataOperationService) GetText(ctx context.Context, req *pb.GetTextRequest) (*pb.GetTextResData, error) {
	content, err := voiceText.ReadText(req.VoiceType.String())
	if err != nil {
		return nil, err
	}
	return &pb.GetTextResData{
		Text: content,
	}, nil
}
