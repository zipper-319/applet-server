package mysql

import (
	"context"
	"gorm.io/gorm"
)

type CloneSpeaker struct {
	//
	Id               int64    `gorm:"column:id;primaryKey" json:"id"`
	SpeakerName      string `gorm:"column:speaker_name" json:"speaker_name"`
	SubmittedSpeaker string `gorm:"column:submitted_speaker" json:"submitted_speaker"`
	Username         string `gorm:"column:username" json:"username"`
	Description      string `gorm:"column:description" json:"description"`
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
	db = db.Select(`*`)
	db = db.Where("username = ?", username)
	db = db.Find(&result)
	return result, db.Error
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

func (m *CloneSpeakerModel) Delete(ctx context.Context, id int64) error {
	db := m.db.WithContext(ctx)
	return db.Where("id = ?", id).Delete(CloneSpeaker{}).Error
}
