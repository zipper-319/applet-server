package biz

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	"applet-server/internal/pkg/util"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type S3UseCase struct {
	*data.Data
	*log.Helper
}

func NewS3UseCase(data *data.Data, logger log.Logger) *S3UseCase {
	return &S3UseCase{Data: data, Helper: log.NewHelper(logger)}
}

func (u *S3UseCase) Upload(ctx context.Context, voiceData []byte, sequence int, username string, voiceType applet.VoiceType) error {

	u.Infof("sequence:%d, username: %s, data length:%d, voiceType:%s", sequence, username, len(voiceData), voiceType)
	key := fmt.Sprintf("%s:%s:%s", util.REDIS_KEY_AWS_S3_USER_Prefix, username, voiceType)
	nextSequence, err := u.RedisClient.Get(ctx, key).Int()
	if err != nil {
		if err == redis.Nil {
			nextSequence = 0
		} else {
			return errors.New("redis error")
		}
	}
	if sequence > nextSequence {
		return errors.New("sequence error")
	}
	// 上传音频到aws s3对象中
	fileName := fmt.Sprintf("%s/%s/%d.wav", username, voiceType.String(), sequence)
	u.Debug("fileName:", fileName)
	url, err := u.Client.Uploading(ctx, voiceData, fileName)
	if err != nil {
		return err
	}
	log.Info("url:", url)

	// 判断是否是重新录制的视频
	if sequence == nextSequence {
		// 修改Redis中的当前进度
		u.RedisClient.Set(ctx, key, sequence+1, 0)
	}
	return err
}

func (u *S3UseCase) Download(ctx context.Context, sequence int, username string, speakerSerial int) ([]byte, error) {
	// 从aws s3中获取音频
	fileName := fmt.Sprintf("%s/%d/%d.pcm", username, speakerSerial, sequence)
	voiceData, err := u.S3.Download(ctx, fileName)
	if err != nil {
		return nil, err
	}
	return voiceData, nil
}
