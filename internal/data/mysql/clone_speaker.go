package mysql

import (
	"context"
	"gorm.io/gorm"
)

type CloneSpeaker struct {
	//
	Id           int64  `gorm:"column:id;primaryKey" json:"id"`
	SpeakerName  string `gorm:"column:speaker_name" json:"speaker_name"`
	SpeakerParam string `gorm:"column:speaker_param" json:"speaker_param"`
	Username     string `gorm:"column:username" json:"username"`
	Description  string `gorm:"column:description" json:"description"`
	IsFinish     bool   `gorm:"column:is_finish" json:"is_finish"`
	*gorm.Model
}

func (CloneSpeaker) TableName() string {
	return "clone_speaker"
}

type CloneSpeakerModel struct {
	db *gorm.DB
}

func NewCloneSpeakerModel(db *gorm.DB) *CloneSpeakerModel {
	return &CloneSpeakerModel{db: db}
}

func (m *CloneSpeakerModel) GetCloneSpeakerList(ctx context.Context, username string) ([]*CloneSpeaker, error) {
	var result []*CloneSpeaker
	var db = m.db.WithContext(ctx)
	db = db.Model(&CloneSpeaker{})
	db = db.Select(`*`)
	db = db.Where("username = ?", username)
	db = db.Find(&result)
	return result, db.Error
}

func (m *CloneSpeakerModel) GetUserSpeakerCount(ctx context.Context, username string) (int64, error) {
	var total int64
	var db = m.db.WithContext(ctx)
	db = db.Model(&CloneSpeaker{})
	db = db.Select(`*`)
	db = db.Where("username = ?", username)
	db = db.Count(&total)
	return total, db.Error
}

func (m *CloneSpeakerModel) Create(ctx context.Context, speaker *CloneSpeaker) error {
	db := m.db.WithContext(ctx)
	return db.Create(speaker).Error
}

func (m *CloneSpeakerModel) Update(ctx context.Context, id int64, speakerName string) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&CloneSpeaker{})
	db = db.Where("id = ?", id)
	return db.UpdateColumn("speaker_name", speakerName).Error
}

func (m *CloneSpeakerModel) SetSpeakerName(ctx context.Context, username, speakerParam, speakerName string) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&CloneSpeaker{})
	db = db.Where("username = ?", username)
	db = db.Where("speaker_param = ?", speakerParam)
	return db.UpdateColumn("speaker_name", speakerName).Error
}

func (m *CloneSpeakerModel) UpdateStatus(ctx context.Context, ids []int64) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&CloneSpeaker{})
	db = db.Where("id in (?)", ids)
	return db.UpdateColumn("is_finish", true).Error
}

func (m *CloneSpeakerModel) Delete(ctx context.Context, id int64) error {
	db := m.db.WithContext(ctx)
	return db.Where("id = ?", id).Delete(&CloneSpeaker{}).Error
}
