package service

import (
	"applet-server/internal/conf"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"

	pb "applet-server/api/v2/applet"
)

type FeedbackService struct {
	pb.UnimplementedFeedbackServer
	addr string
}

type CollectQAReq struct {
	Agentid  int32  `json:"agentid"`
	Lang     string `json:"lang"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Channel  string `json:"channel"`
}

type CollectQAResp struct {
	Code    int  `json:"code"`
	Status  bool `json:"status"`
	Agentid int  `json:"agentid"`
	QaId    int  `json:"qaId"`
}

func NewFeedbackService(app *conf.App) *FeedbackService {
	return &FeedbackService{addr: app.Feedback.Addr}
}

func (s *FeedbackService) Collect(ctx context.Context, req *pb.CollectReq) (*emptypb.Empty, error) {
	if req.QaType == pb.QAType_CommonQA{
		req.AgentId = 0
	}
	qaReq := &CollectQAReq{
		Agentid:  req.AgentId,
		Lang:     req.Language,
		Question: req.Question,
		Answer:   req.Answer,
		Channel:  "weixin",
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
