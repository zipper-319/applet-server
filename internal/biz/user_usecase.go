package biz

import (
	"applet-server/internal/data"
	"applet-server/internal/data/mysql"
	"applet-server/internal/pkg/log"
	"context"
)

type UserUseCase struct {
	*data.Data
	*log.MyLogger
}

func NewUserUseCase(data *data.Data, logger *log.MyLogger) *UserUseCase {
	return &UserUseCase{Data: data, MyLogger: logger}
}

func (u *UserUseCase) GetUserByNameAndPhone(ctx context.Context, name string, phone string) (*mysql.User, error) {
	repo := mysql.NewUsersModel(u.DB)
	user, err := repo.GetUser(ctx, name, phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
