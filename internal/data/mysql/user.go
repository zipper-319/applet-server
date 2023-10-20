package mysql

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	//
	Id int64 `json:"id" gorm:"primary_key" gorm:"column:id"`
	// 微信openID
	OpendId string `json:"openId" gorm:"column:openId"`
	//  姓名
	FullName string `json:"FullName" gorm:"column:full_name"`
	// 电话号码
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number"`
	// 角色
	Role int `json:"role" gorm:"column:role"`
	// 账户状态=={‘0’:’暂停’,’1’:’启用’}
	Status int `json:"status" gorm:"column:status"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func (User) TableName() string {
	return "user"
}

type UsersModel struct {
	db *gorm.DB
}

func NewUsersModel(db *gorm.DB) *UsersModel {
	return &UsersModel{db: db}
}

func (m *UsersModel) GetUser(ctx context.Context, fullName, phoneNumber string) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`id`)
	db = db.Where("full_name = ? and phone_number= ?", fullName, phoneNumber)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}
