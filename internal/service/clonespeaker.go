package service

import (
	"applet-server/internal/biz"
	jwtUtil "applet-server/internal/pkg/jwt"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "applet-server/api/v2/applet"
)

type CloneSpeakerService struct {
	pb.UnimplementedCloneSpeakerServer
	uc *biz.CloneSpeakerUseCase
}

func NewCloneSpeakerService(useCase *biz.CloneSpeakerUseCase) *CloneSpeakerService {
	return &CloneSpeakerService{uc:useCase}
}

func (s *CloneSpeakerService) GetCloneSpeaker(ctx context.Context, req *emptypb.Empty) (*pb.GetCloneSpeakerResult, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	return s.uc.GetSpeakerList(ctx, tokenInfo.Username)
}
func (s *CloneSpeakerService) UpdateCloneSpeaker(ctx context.Context, req *pb.UpdateCloneSpeakerRequest) (*emptypb.Empty, error) {
	if err := s.uc.UpdateSpeaker(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (s *CloneSpeakerService) DelCloneSpeaker(ctx context.Context, req *pb.DelCloneSpeakerRequest) (*emptypb.Empty, error) {
	if err := s.uc.DelSpeaker(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
