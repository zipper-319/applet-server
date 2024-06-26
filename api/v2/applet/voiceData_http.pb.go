// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.6.1
// source: v2/applet/voiceData.proto

package applet

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationVoiceDataOperationcommit = "/applet.v2.VoiceDataOperation/commit"
const OperationVoiceDataOperationdownloadVoice = "/applet.v2.VoiceDataOperation/downloadVoice"
const OperationVoiceDataOperationgetProgress = "/applet.v2.VoiceDataOperation/getProgress"
const OperationVoiceDataOperationgetText = "/applet.v2.VoiceDataOperation/getText"
const OperationVoiceDataOperationputVoiceData = "/applet.v2.VoiceDataOperation/putVoiceData"

type VoiceDataOperationHTTPServer interface {
	// Commit 提交已完成录制音频
	Commit(context.Context, *CommitRequest) (*CommitResData, error)
	// DownloadVoice 下载已经录制音频
	DownloadVoice(context.Context, *DownloadReqData) (*DownloadResData, error)
	// GetProgress 获取音频录制进度，返回已录制的音频数量
	GetProgress(context.Context, *ProgressRequest) (*ProgressResData, error)
	// GetText 获取录音文本
	GetText(context.Context, *GetTextRequest) (*GetTextResData, error)
	// PutVoiceData 训练数据上传
	PutVoiceData(context.Context, *VoiceDataReqData) (*VoiceDataResData, error)
}

func RegisterVoiceDataOperationHTTPServer(s *http.Server, srv VoiceDataOperationHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v2/voice-data/video/upload", _VoiceDataOperation_PutVoiceData0_HTTP_Handler(srv))
	r.GET("/api/v2/voice-data/get-progress", _VoiceDataOperation_GetProgress0_HTTP_Handler(srv))
	r.GET("/api/v2/voice-data/download", _VoiceDataOperation_DownloadVoice0_HTTP_Handler(srv))
	r.POST("/api/v2/voice-data/commit", _VoiceDataOperation_Commit0_HTTP_Handler(srv))
	r.GET("/api/v2/voice-data/text/get", _VoiceDataOperation_GetText0_HTTP_Handler(srv))
}

func _VoiceDataOperation_PutVoiceData0_HTTP_Handler(srv VoiceDataOperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VoiceDataReqData
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVoiceDataOperationputVoiceData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PutVoiceData(ctx, req.(*VoiceDataReqData))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VoiceDataResData)
		return ctx.Result(200, reply)
	}
}

func _VoiceDataOperation_GetProgress0_HTTP_Handler(srv VoiceDataOperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProgressRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVoiceDataOperationgetProgress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProgress(ctx, req.(*ProgressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProgressResData)
		return ctx.Result(200, reply)
	}
}

func _VoiceDataOperation_DownloadVoice0_HTTP_Handler(srv VoiceDataOperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadReqData
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVoiceDataOperationdownloadVoice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DownloadVoice(ctx, req.(*DownloadReqData))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadResData)
		return ctx.Result(200, reply)
	}
}

func _VoiceDataOperation_Commit0_HTTP_Handler(srv VoiceDataOperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommitRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVoiceDataOperationcommit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Commit(ctx, req.(*CommitRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommitResData)
		return ctx.Result(200, reply)
	}
}

func _VoiceDataOperation_GetText0_HTTP_Handler(srv VoiceDataOperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTextRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVoiceDataOperationgetText)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetText(ctx, req.(*GetTextRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTextResData)
		return ctx.Result(200, reply)
	}
}

type VoiceDataOperationHTTPClient interface {
	Commit(ctx context.Context, req *CommitRequest, opts ...http.CallOption) (rsp *CommitResData, err error)
	DownloadVoice(ctx context.Context, req *DownloadReqData, opts ...http.CallOption) (rsp *DownloadResData, err error)
	GetProgress(ctx context.Context, req *ProgressRequest, opts ...http.CallOption) (rsp *ProgressResData, err error)
	GetText(ctx context.Context, req *GetTextRequest, opts ...http.CallOption) (rsp *GetTextResData, err error)
	PutVoiceData(ctx context.Context, req *VoiceDataReqData, opts ...http.CallOption) (rsp *VoiceDataResData, err error)
}

type VoiceDataOperationHTTPClientImpl struct {
	cc *http.Client
}

func NewVoiceDataOperationHTTPClient(client *http.Client) VoiceDataOperationHTTPClient {
	return &VoiceDataOperationHTTPClientImpl{client}
}

func (c *VoiceDataOperationHTTPClientImpl) Commit(ctx context.Context, in *CommitRequest, opts ...http.CallOption) (*CommitResData, error) {
	var out CommitResData
	pattern := "/api/v2/voice-data/commit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationVoiceDataOperationcommit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VoiceDataOperationHTTPClientImpl) DownloadVoice(ctx context.Context, in *DownloadReqData, opts ...http.CallOption) (*DownloadResData, error) {
	var out DownloadResData
	pattern := "/api/v2/voice-data/download"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVoiceDataOperationdownloadVoice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VoiceDataOperationHTTPClientImpl) GetProgress(ctx context.Context, in *ProgressRequest, opts ...http.CallOption) (*ProgressResData, error) {
	var out ProgressResData
	pattern := "/api/v2/voice-data/get-progress"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVoiceDataOperationgetProgress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VoiceDataOperationHTTPClientImpl) GetText(ctx context.Context, in *GetTextRequest, opts ...http.CallOption) (*GetTextResData, error) {
	var out GetTextResData
	pattern := "/api/v2/voice-data/text/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationVoiceDataOperationgetText))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VoiceDataOperationHTTPClientImpl) PutVoiceData(ctx context.Context, in *VoiceDataReqData, opts ...http.CallOption) (*VoiceDataResData, error) {
	var out VoiceDataResData
	pattern := "/api/v2/voice-data/video/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationVoiceDataOperationputVoiceData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
