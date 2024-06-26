// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: v2/applet/tts.proto

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
	TTSService_GetTTSConfig_FullMethodName = "/applet.v2.TTSService/GetTTSConfig"
)

// TTSServiceClient is the client API for TTSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TTSServiceClient interface {
	GetTTSConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTTSConfigResult, error)
}

type tTSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTTSServiceClient(cc grpc.ClientConnInterface) TTSServiceClient {
	return &tTSServiceClient{cc}
}

func (c *tTSServiceClient) GetTTSConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTTSConfigResult, error) {
	out := new(GetTTSConfigResult)
	err := c.cc.Invoke(ctx, TTSService_GetTTSConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TTSServiceServer is the server API for TTSService service.
// All implementations must embed UnimplementedTTSServiceServer
// for forward compatibility
type TTSServiceServer interface {
	GetTTSConfig(context.Context, *emptypb.Empty) (*GetTTSConfigResult, error)
	mustEmbedUnimplementedTTSServiceServer()
}

// UnimplementedTTSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTTSServiceServer struct {
}

func (UnimplementedTTSServiceServer) GetTTSConfig(context.Context, *emptypb.Empty) (*GetTTSConfigResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTTSConfig not implemented")
}
func (UnimplementedTTSServiceServer) mustEmbedUnimplementedTTSServiceServer() {}

// UnsafeTTSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TTSServiceServer will
// result in compilation errors.
type UnsafeTTSServiceServer interface {
	mustEmbedUnimplementedTTSServiceServer()
}

func RegisterTTSServiceServer(s grpc.ServiceRegistrar, srv TTSServiceServer) {
	s.RegisterService(&TTSService_ServiceDesc, srv)
}

func _TTSService_GetTTSConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TTSServiceServer).GetTTSConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TTSService_GetTTSConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TTSServiceServer).GetTTSConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TTSService_ServiceDesc is the grpc.ServiceDesc for TTSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TTSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "applet.v2.TTSService",
	HandlerType: (*TTSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTTSConfig",
			Handler:    _TTSService_GetTTSConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2/applet/tts.proto",
}
