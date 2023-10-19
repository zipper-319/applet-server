package service

import (
	"applet-server/internal/biz"
	"context"
	"encoding/base64"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "applet-server/api/v2/applet"
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
func (s *VoiceDataOperationService) GetProgress(ctx context.Context, req *emptypb.Empty) (*pb.ProgressResData, error) {

	return &pb.ProgressResData{
		CurrentNumber: 0,
		FinishedTime:  "",
		AwaitTrain:    0,
	}, nil
}
func (s *VoiceDataOperationService) DownloadVoice(ctx context.Context, req *pb.DownloadReqData) (*pb.DownloadResData, error) {
	return &pb.DownloadResData{}, nil
}
func (s *VoiceDataOperationService) Commit(ctx context.Context, req *emptypb.Empty) (*pb.CommitResData, error) {
	return &pb.CommitResData{}, nil
}
func (s *VoiceDataOperationService) GetText(ctx context.Context, req *emptypb.Empty) (*pb.GetTextResData, error) {
	return &pb.GetTextResData{}, nil
}
