package service

import (
	"applet-server/internal/conf"
	"applet-server/internal/pkg/http"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/log"
	"time"

	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "applet-server/api/v2/applet"
)

type FeedbackService struct {
	pb.UnimplementedFeedbackServer
	addr string
	*log.MyLogger
}

type CollectQAReq struct {
	AgentId  int32  `json:"agentid"`
	Lang     string `json:"lang"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

type CollectQAResp struct {
	Code    int  `json:"code"`
	Status  bool `json:"status"`
	AgentId int  `json:"agentid"`
	QaId    int  `json:"qaId"`
}

type CollectResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type CollectBugReq struct {
	DislikeCommon DislikeCommon `json:"dislikecommon"`
	Mark          string        `json:"mark"`
}

type CollectUseCaseReq struct {
	LikeCommon   LikeCommon    `json:"likeCommon"`
	Mark         string        `json:"mark"`
	Source       string        `json:"source"`
	Domain       string        `json:"domain"`
	Intent       string        `json:"intentname"`
	IntentParams []IntentParam `json:"intentparams"`
	QGroupId     string        `json:"qgroupid"`
}

type IntentParam struct {
	SlotType  string `json:"slottype"`
	SlotValue string `json:"slotvalue"`
	SlotName  string `json:"slotname"`
}

type DislikeCommon struct {
	AgentId        int32  `json:"agentid"`
	QuestionId     string `json:"questionid"`
	Question       string `json:"question"`
	Answer         string `json:"answer"`
	BugType        string `json:"bugtype"`
	FeedbackTime   string `json:"feedbacktime"`
	FeedbackPerson string `json:"feedbackperson"`
	Env            string `json:"env"`
	Lang           string `json:"lang"`
	SessionId      string `json:"sessionid"`
	RobotType      string `json:"robottype"`
}

type LikeCommon struct {
	AgentId        int32  `json:"agentid"`
	QuestionId     string `json:"questionid"`
	Question       string `json:"question"`
	Answer         string `json:"answer"`
	BugType        string `json:"bugtype"`
	FeedbackTime   string `json:"feedbacktime"`
	FeedbackPerson string `json:"feedbackperson"`
	Env            string `json:"env"`
	Lang           string `json:"lang"`
	SessionId      string `json:"sessionid"`
	RobotType      string `json:"robottype"`
}

func NewFeedbackService(app *conf.App, logger *log.MyLogger) *FeedbackService {
	return &FeedbackService{addr: app.Feedback.Addr, MyLogger: logger}
}

func (s *FeedbackService) Collect(ctx context.Context, req *pb.CollectReq) (*emptypb.Empty, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}
	s.Infof("tokenInfo: %+v", tokenInfo)
	username := tokenInfo.Username
	if req.QaType == pb.QAType_CommonQA {
		req.AgentId = 0
	}
	qaReq := &CollectQAReq{
		AgentId:  req.AgentId,
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
	addr := fmt.Sprintf("http://%s%s", s.addr, "/v2/ux/docqa/markqa")
	result, err := http.Post(addr, qaReqStr)
	if err != nil {
		return nil, err
	}
	s.Infof("result: %s", result)
	var respData CollectQAResp
	if err := json.Unmarshal(result, &respData); err != nil {
		return nil, err
	}
	if respData.Code == 0 && respData.Status {
		return &emptypb.Empty{}, nil
	} else {
		return &emptypb.Empty{}, errors.New("feedback failed")
	}
}

func (s *FeedbackService) CollectLike(ctx context.Context, req *pb.CollectLikeReq) (*emptypb.Empty, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}

	username := tokenInfo.Username
	s.Infof("tokenInfo: %+v; username:%s", tokenInfo, username)
	addr := fmt.Sprintf("http://%s%s", s.addr, "/v2/ux/docqa/like")
	intentParams := make([]IntentParam, 0, len(req.Entities))
	for _, entity := range req.Entities {
		intentParams = append(intentParams, IntentParam{
			SlotType:  entity.Type,
			SlotValue: entity.Value,
			SlotName:  entity.Name,
		})
	}
	collectReq := &CollectUseCaseReq{
		LikeCommon: LikeCommon{
			AgentId:        req.AgentId,
			QuestionId:     req.QuestionId,
			Question:       req.Question,
			Answer:         req.Answer,
			FeedbackTime:   time.Now().Format(time.RFC3339),
			FeedbackPerson: username,
			Env:            req.EnvType,
			Lang:           req.Language,
			RobotType:      "weixin",
		},
		Mark:         "",
		Source:       req.Source,
		Domain:       req.Domain,
		Intent:       req.Intent,
		IntentParams: intentParams,
		QGroupId:     "",
	}
	collectReqStr, _ := json.Marshal(collectReq)
	result, err := http.Post(addr, collectReqStr)
	if err != nil {
		return nil, err
	}
	s.Infof("result: %s", result)
	var resp CollectResp
	if err := json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}
	if resp.Code == 0 && resp.Status {
		return &emptypb.Empty{}, nil
	} else {
		return nil, errors.New(resp.Message)
	}

}
func (s *FeedbackService) CollectDislike(ctx context.Context, req *pb.CollectDislikeReq) (*emptypb.Empty, error) {
	tokenInfo, ok := jwtUtil.GetTokenInfo(ctx)
	if !ok {
		return nil, jwtUtil.ErrTokenInvalid
	}

	username := tokenInfo.Username
	s.Infof("tokenInfo: %+v; username:%s", tokenInfo, username)
	addr := fmt.Sprintf("http://%s%s", s.addr, "/v2/ux/docqa/dislike")
	collectReq := &CollectBugReq{
		DislikeCommon: DislikeCommon{
			AgentId:        req.AgentId,
			QuestionId:     req.QuestionId,
			Question:       req.Question,
			Answer:         req.Answer,
			BugType:        fmt.Sprintf("%s-%s", req.BugType.String(), req.GetBugDesc()),
			FeedbackTime:   time.Now().Format(time.RFC3339),
			FeedbackPerson: username,
			Env:            req.EnvType,
			Lang:           req.Language,
			SessionId:      req.SessionId,
			RobotType:      "weixin",
		},
		Mark: fmt.Sprintf("%s-%s", req.Reality, req.Expectation),
	}
	collectReqStr, _ := json.Marshal(collectReq)
	result, err := http.Post(addr, collectReqStr)
	if err != nil {
		return nil, err
	}
	s.Infof("result: %s", result)
	var resp CollectResp
	json.Unmarshal(result, &resp)
	if resp.Code == 0 && resp.Status {
		return &emptypb.Empty{}, nil
	} else {
		return nil, errors.New(resp.Message)
	}
}
