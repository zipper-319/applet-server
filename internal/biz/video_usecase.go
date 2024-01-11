package biz

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	"applet-server/internal/data/mysql"
	"applet-server/internal/pkg/log"
	"applet-server/internal/pkg/util"
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type VideoUseCase struct {
	*data.Data
	*log.MyLogger
}

func NewVideoUseCase(data *data.Data, logger *log.MyLogger) *VideoUseCase {
	return &VideoUseCase{Data: data, MyLogger: logger}
}

func (u *VideoUseCase) Upload(ctx context.Context, voiceData []byte, sequence int, username string, voiceType applet.VoiceType) error {
	var speakerParam string
	flag := applet.Flag_continue
	u.Infof("sequence:%d, username: %s, data length:%d, voiceType:%s", sequence, username, len(voiceData), voiceType)
	key := fmt.Sprintf("%s:%s:%s", util.REDIS_KEY_AWS_S3_USER_Prefix, username, voiceType)
	nextSequence, err := u.RedisClient.HGet(ctx, key, "sequence").Int()
	if err != nil {
		if err == redis.Nil {
			nextSequence = 0
			flag = applet.Flag_start
		} else {
			return errors.New("redis error")
		}
	}
	if sequence > nextSequence {
		return errors.New("sequence error")
	}
	speakerParam, err = u.RedisClient.HGet(ctx, key, "speaker").Result()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("HGet;redis error:%v", err)
	}
	// 创建发音人Id
	if speakerParam == "" {
		speakerRepo := mysql.NewCloneSpeakerModel(u.DB)
		count, err := speakerRepo.GetUserSpeakerCount(ctx, username)
		if err != nil {
			return err
		}
		speakerParam = fmt.Sprintf("%s%s%d", username, voiceType, count+1)
		u.RedisClient.HSet(ctx, key, "speaker", speakerParam)
	}
	// 上传音频到aws s3对象中
	fileName := fmt.Sprintf("%s_%s_%d.wav", username, voiceType.String(), sequence)
	u.Debugf("fileName:%s, speakerParam:%s", fileName, speakerParam)
	videoDataReader := bytes.NewReader(voiceData)
	if err := u.Train.SaveVideo(videoDataReader, username, speakerParam, fileName, flag); err != nil {
		return err
	}

	// 判断是否是重新录制的视频
	if sequence == nextSequence {
		// 修改Redis中的当前进度
		u.RedisClient.HSet(ctx, key, "sequence", sequence+1)
	}
	return nil
}

func (u *VideoUseCase) Download(ctx context.Context, sequence int, username string, speakerSerial int) ([]byte, error) {
	// 从aws s3中获取音频
	fileName := fmt.Sprintf("%s/%d/%d.pcm", username, speakerSerial, sequence)
	voiceData, err := u.S3.Download(ctx, fileName)
	if err != nil {
		return nil, err
	}
	return voiceData, nil
}

func (u *VideoUseCase) Commit(ctx context.Context, username, speakerParam string) error {
	if err := u.Train.SaveVideo(nil, username, speakerParam, "", applet.Flag_end); err != nil {
		return err
	}
	return nil
}
