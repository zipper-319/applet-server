package service

import (
	pb "applet-server/api/v2/applet"
	"applet-server/internal/biz"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FeedbackService struct {
	pb.UnimplementedFeedbackServer
	addr string

	*biz.FeedBackUseCase
}

func NewFeedbackService(useCase *biz.FeedBackUseCase) *FeedbackService {
	return &FeedbackService{FeedBackUseCase: useCase}
}

func (s *FeedbackService) Collect(ctx context.Context, req *pb.CollectReq) (*emptypb.Empty, error) {
	return s.FeedBackUseCase.Collect(ctx, req)
}

func (s *FeedbackService) CollectLike(ctx context.Context, req *pb.CollectLikeReq) (*emptypb.Empty, error) {
	return s.FeedBackUseCase.CollectLike(ctx, req)
}

func (s *FeedbackService) CollectDislike(ctx context.Context, req *pb.CollectDislikeReq) (*emptypb.Empty, error) {
	return s.FeedBackUseCase.CollectDislike(ctx, req)
}
