package service

import (
	pb "applet-server/api/v2/applet"
	"applet-server/internal/biz"
	jwtUtil "applet-server/internal/pkg/jwt"
	"context"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	*biz.UserUseCase
}

func NewAccountService(useCase *biz.UserUseCase) *AccountService {
	return &AccountService{
		UserUseCase: useCase,
	}
}

func (s *AccountService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	account, err := s.GetUserByNameAndPhone(ctx, req.FullName, req.PhoneNumber)
	if err != nil {
		return nil, err
	}
	token, err := jwtUtil.GetToken(req.GetFullName(), req.PhoneNumber, account.Role)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Token: token,
	}, nil
}
