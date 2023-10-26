package service

import (
	pb "applet-server/api/v2/applet"
	"applet-server/internal/biz"
	"applet-server/internal/data"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/util"
	"applet-server/internal/pkg/voiceText"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
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
	if err := s.S3UseCase.Upload(ctx, voiceData, int(req.Sequence), username, req.VoiceType); err != nil {
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

	progressKey := fmt.Sprintf("%s:%s:%s", util.REDIS_KEY_AWS_S3_USER_Prefix, username, req.VoiceType)
	sequence, err := s.S3UseCase.Data.RedisClient.Get(ctx, progressKey).Int64()

	finishedKey := fmt.Sprintf("finishedTime:%s:%s", username, req.VoiceType)
	finishedTime, err := s.S3UseCase.Data.RedisClient.Get(ctx, finishedKey).Int64()
	if err == redis.Nil {
		finishedTime = 0
	} else {
		return nil, errors.New("redis error")
	}
	return &pb.ProgressResData{
		CurrentNumber: int32(sequence),
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
		s3NameList = append(s3NameList, fmt.Sprintf("%s/%d.pcm", username, i))
	}
	key := fmt.Sprintf("finishedTime:%s:%s", username, req.Speaker)

	if s.S3UseCase.Data.RedisClient.SetNX(ctx, key, time.Now().Unix(), 0).Val() {
		progressKey := fmt.Sprintf("%s:%s:%s", util.REDIS_KEY_AWS_S3_USER_Prefix, username, req.VoiceType)
		s.Data.RedisClient.Del(ctx, progressKey)
		return &pb.CommitResData{}, nil
	} else {
		return nil, errors.New("speaker has existed")
	}
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

func (s *VoiceDataOperationService) HandlerFormData(ctx context.Context, file *data.FileObject, req *pb.UploadFilesRequest) (*emptypb.Empty, error) {
	s.Info(req)
	s.Info(file)
	return &emptypb.Empty{}, nil
}
