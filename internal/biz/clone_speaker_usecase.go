package biz

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	"applet-server/internal/data/mysql"
	"applet-server/internal/pkg/util"
	"context"
)

type CloneSpeakerUseCase struct {
	*data.Data
	repo *mysql.CloneSpeakerModel
}

func NewCloneSpeakerUseCase(data *data.Data) *CloneSpeakerUseCase {
	return &CloneSpeakerUseCase{Data: data, repo: mysql.NewCloneSpeakerModel(data.DB)}
}

func (uc *CloneSpeakerUseCase) GetSpeakerList(ctx context.Context, username string) (*applet.GetCloneSpeakerResult, error) {
	result, err := uc.repo.GetCloneSpeakerList(ctx, username)
	if err != nil {
		return nil, err
	}
	speakerList := make([]*applet.GetCloneSpeakerResult_CloneSpeaker, 0, len(result))
	for _, v := range result {
		speakerList = append(speakerList, &applet.GetCloneSpeakerResult_CloneSpeaker{
			Id:          v.Id,
			SpeakerName: v.SpeakerName,
			IsFinish:    false,
			Description: v.Description,
			CreateTime:  v.CreatedAt.Unix(),
			UpdateTime:  v.UpdatedAt.Unix(),
		})
	}
	return &applet.GetCloneSpeakerResult{
		CloneSpeakerList: speakerList,
		TrainTime:        util.DefaultTrainDuration,
	}, nil
}

func (uc *CloneSpeakerUseCase) UpdateSpeaker(ctx context.Context, req *applet.UpdateCloneSpeakerRequest) error {
	return uc.repo.Update(ctx, req.Id, req.SpeakerName)
}

func (uc *CloneSpeakerUseCase) DelSpeaker(ctx context.Context, req *applet.DelCloneSpeakerRequest) error {
	return uc.repo.Delete(ctx, req.Id)
}

func (uc *CloneSpeakerUseCase) CreateSpeaker(ctx context.Context, speaker, username string) error {
	return uc.repo.Create(ctx, &mysql.CloneSpeaker{
		SpeakerName:      speaker,
		SubmittedSpeaker: speaker,
		Username:         username,
	})
}
