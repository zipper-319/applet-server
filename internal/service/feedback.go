package service

import (
	"applet-server/internal/conf"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/log"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"

	pb "applet-server/api/v2/applet"
)

type FeedbackService struct {
	pb.UnimplementedFeedbackServer
	addr string
	*log.MyLogger
}

type CollectQAReq struct {
	Agentid  int32  `json:"agentid"`
	Lang     string `json:"lang"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

type CollectQAResp struct {
	Code    int  `json:"code"`
	Status  bool `json:"status"`
	Agentid int  `json:"agentid"`
	QaId    int  `json:"qaId"`
}

func NewFeedbackService(app *conf.App, logger *log.MyLogger) *FeedbackService {
	return &FeedbackService{addr: app.Feedback.Addr,  MyLogger: logger}
}

func (s *FeedbackService) Collect(ctx context.Context, req *pb.CollectReq) (*emptypb.Empty, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	s.Infof("tokenInfo: %+v", tokenInfo)
	username := tokenInfo.Username
	if req.QaType == pb.QAType_CommonQA{
		req.AgentId = 0
	}
	qaReq := &CollectQAReq{
		Agentid:  req.AgentId,
		Lang:     req.Language,
		Question: req.Question,
		Answer:   req.Answer,
		Channel:  "weixin",
		Username: username,
	}

	qaReqStr, err := json.Marshal(qaReq)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s%s", s.addr, "/v2/ux/docqa/markqa"), "application/json", bytes.NewReader(qaReqStr))
	if err != nil {
		return nil, err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Info(string(result))
	var respData CollectQAResp
	json.Unmarshal(result, &respData)
	if respData.Code == 0 && respData.Status {
		return &emptypb.Empty{}, nil
	} else {
		return &emptypb.Empty{}, errors.New("feedback failed")
	}
}
