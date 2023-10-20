package biz

import (
	"applet-server/internal/data"
	"applet-server/internal/data/mysql"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	*data.Data
	*log.Helper
}

func NewUserUseCase(data *data.Data, logger log.Logger) *UserUseCase {
	return &UserUseCase{Data: data, Helper: log.NewHelper(logger)}
}

func (u *UserUseCase) GetUserByNameAndPhone(ctx context.Context, name string, phone string) (*mysql.User, error) {
	repo := mysql.NewUsersModel(u.DB)
	user, err := repo.GetUser(ctx, name, phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
