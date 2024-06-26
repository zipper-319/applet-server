package biz

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/biz/tts"
	"applet-server/internal/data"
	"applet-server/internal/data/mysql"
	"applet-server/internal/pkg/util"
	"context"
)

type CloneSpeakerUseCase struct {
	*data.Data
	repo      *mysql.CloneSpeakerModel
	ttsClient *tts.TTSClient
}

func NewCloneSpeakerUseCase(data *data.Data, client *tts.TTSClient) *CloneSpeakerUseCase {
	return &CloneSpeakerUseCase{Data: data, repo: mysql.NewCloneSpeakerModel(data.DB), ttsClient: client}
}

func (uc *CloneSpeakerUseCase) GetSpeakerList(ctx context.Context, username string) (*applet.GetCloneSpeakerResult, error) {
	result, err := uc.repo.GetCloneSpeakerList(ctx, username)
	if err != nil {
		return nil, err
	}
	userSpaceSpeakers, err := uc.ttsClient.GetUserSpeaker(ctx, username)
	if err != nil {
		return nil, err
	}
	finishedSpeakerIdList := make([]int64, 0, len(result))

	speakerList := make([]*applet.GetCloneSpeakerResult_CloneSpeaker, 0, len(result))
	for _, v := range result {
		status := v.IsFinish
		if !status {
			status = util.IsSpeakerExist(userSpaceSpeakers, v.SpeakerParam)
			if status{
				finishedSpeakerIdList = append(finishedSpeakerIdList, v.Id)
			}
		}
		speakerList = append(speakerList, &applet.GetCloneSpeakerResult_CloneSpeaker{
			Id:           v.Id,
			SpeakerName:  v.SpeakerName,
			SpeakerParam: v.SpeakerParam,
			IsFinish:     status,
			Description:  v.Description,
			CreateTime:   v.CreatedAt.Unix(),
			UpdateTime:   v.UpdatedAt.Unix(),
		})
	}
	if len(finishedSpeakerIdList) != 0 {
		uc.repo.UpdateStatus(ctx, finishedSpeakerIdList)
	}
	return &applet.GetCloneSpeakerResult{
		CloneSpeakerList: speakerList,
		TrainTime:        util.DefaultTrainDuration,
	}, nil
}

func (uc *CloneSpeakerUseCase) UpdateSpeaker(ctx context.Context, req *applet.UpdateCloneSpeakerRequest) error {
	return uc.repo.Update(ctx, req.Id, req.SpeakerName)
}

func (uc *CloneSpeakerUseCase) CommitSpeakerName(ctx context.Context, username, speakerParam, speakerName string) error {
	return uc.repo.SetSpeakerName(ctx, username, speakerParam, speakerName)
}

func (uc *CloneSpeakerUseCase) DelSpeaker(ctx context.Context, req *applet.DelCloneSpeakerRequest) error {
	return uc.repo.Delete(ctx, req.Id)
}

func (uc *CloneSpeakerUseCase) CreateSpeaker(ctx context.Context, username, speakerName, speakerParam string) error {
	return uc.repo.Create(ctx, &mysql.CloneSpeaker{
		SpeakerName:  speakerName,
		SpeakerParam: speakerParam,
		Username:     username,
	})
}

func (uc *CloneSpeakerUseCase) GetSpeakerMap(ctx context.Context, username string) (map[string]string, error) {
	speakerList, err := uc.repo.GetCloneSpeakerList(ctx, username)
	if err != nil {
		return nil, err
	}
	speakerMap := make(map[string]string, len(speakerList))
	for _, v := range speakerList {
		speakerMap[v.SpeakerParam] = v.SpeakerName
	}
	return speakerMap, nil
}
