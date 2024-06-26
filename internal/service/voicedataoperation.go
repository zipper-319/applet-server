package service

import (
	pb "applet-server/api/v2/applet"
	"applet-server/internal/biz"
	"applet-server/internal/data"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/log"
	"applet-server/internal/pkg/util"
	"applet-server/internal/pkg/voiceText"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"time"
)

type VoiceDataOperationService struct {
	pb.UnimplementedVoiceDataOperationServer
	uc      *biz.VideoUseCase
	speaker *biz.CloneSpeakerUseCase
	*log.MyLogger
}

func NewVoiceDataOperationService(useCase *biz.VideoUseCase, speakerUseCase *biz.CloneSpeakerUseCase, logger *log.MyLogger) *VoiceDataOperationService {
	return &VoiceDataOperationService{
		uc:       useCase,
		speaker:  speakerUseCase,
		MyLogger: logger,
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
	if err := s.uc.Upload(ctx, voiceData, int(req.Sequence), username, req.VoiceType); err != nil {
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
	sequence, err := s.uc.Data.RedisClient.HGet(ctx, progressKey, "sequence").Int64()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("HGet sequence failed:%v", err)
	}
	//finishedKey := fmt.Sprintf("finishedTime:%s:%s", username, req.VoiceType)
	//finishedTime, err := s.uc.Data.RedisClient.Get(ctx, finishedKey).Int64()
	//if err != nil {
	//	if err == redis.Nil {
	//		finishedTime = 0
	//	} else {
	//		return nil, fmt.Errorf("get finishedTime failed:%v", err)
	//	}
	//}

	return &pb.ProgressResData{
		CurrentNumber: int32(sequence),
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
	//s3NameList := make([]string, 0, 50)
	//for i := 0; i < voiceText.VoiceDataSize[req.VoiceType]; i++ {
	//	s3NameList = append(s3NameList, fmt.Sprintf("%s/%d.pcm", username, i))
	//}
	progressKey := fmt.Sprintf("%s:%s:%s", util.REDIS_KEY_AWS_S3_USER_Prefix, username, req.VoiceType)
	speakerParam, err := s.uc.RedisClient.HGet(ctx, progressKey, "speaker").Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("get speaker failed:%v", err)
	}

	key := fmt.Sprintf("finishedTime:%s:%s", username, speakerParam)

	if s.uc.RedisClient.SetNX(ctx, key, time.Now().Unix(), 0).Val() {

		if err := s.uc.Commit(ctx, username, speakerParam); err != nil {
			return nil, err
		}
		if err := s.speaker.CreateSpeaker(ctx, username, req.Speaker, speakerParam); err != nil {
			return nil, err
		}

		s.uc.RedisClient.Del(ctx, progressKey)

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
	//uploadMinio(file.File, file.FileName)
	speakerParam, err := s.uc.UploadFiles(ctx, file.File, file.Username, file.FileName)
	if err != nil {
		return nil, err
	}

	if err := s.speaker.CreateSpeaker(ctx, file.Username, req.Speaker, speakerParam); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func uploadMinio(file io.Reader, filename string) error {
	endpoint := "10.12.32.96:9100"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Info(err)
	}

	ctx := context.Background()
	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	location := "us-east-1"

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {

		return err
	}
	if !exists {

		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return err

		}
	}

	// Upload the zip file
	//objectName := "video/data机器人.mp4"
	//filePath := "/home/data/下载/video/data机器人.mp4"
	contentType := "video/mp4"

	// Upload the zip file with FPutObject
	info, err := minioClient.PutObject(ctx, bucketName, filename, file, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Info(err)
	}

	log.Info("Successfully uploaded %s of size %s\n", filename, info.Location)
	return nil
}
