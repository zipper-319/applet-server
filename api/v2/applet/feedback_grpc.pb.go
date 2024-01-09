// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: v2/applet/feedback.proto

package applet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Feedback_Collect_FullMethodName = "/applet.v2.Feedback/Collect"
)

// FeedbackClient is the client API for Feedback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackClient interface {
	Collect(ctx context.Context, in *CollectReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type feedbackClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackClient(cc grpc.ClientConnInterface) FeedbackClient {
	return &feedbackClient{cc}
}

func (c *feedbackClient) Collect(ctx context.Context, in *CollectReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Feedback_Collect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServer is the server API for Feedback service.
// All implementations must embed UnimplementedFeedbackServer
// for forward compatibility
type FeedbackServer interface {
	Collect(context.Context, *CollectReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedFeedbackServer()
}

// UnimplementedFeedbackServer must be embedded to have forward compatible implementations.
type UnimplementedFeedbackServer struct {
}

func (UnimplementedFeedbackServer) Collect(context.Context, *CollectReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedFeedbackServer) mustEmbedUnimplementedFeedbackServer() {}

// UnsafeFeedbackServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackServer will
// result in compilation errors.
type UnsafeFeedbackServer interface {
	mustEmbedUnimplementedFeedbackServer()
}

func RegisterFeedbackServer(s grpc.ServiceRegistrar, srv FeedbackServer) {
	s.RegisterService(&Feedback_ServiceDesc, srv)
}

func _Feedback_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Feedback_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServer).Collect(ctx, req.(*CollectReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Feedback_ServiceDesc is the grpc.ServiceDesc for Feedback service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feedback_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "applet.v2.Feedback",
	HandlerType: (*FeedbackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collect",
			Handler:    _Feedback_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2/applet/feedback.proto",
}
