// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: notice.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NoticeServiceClient is the client API for NoticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoticeServiceClient interface {
	//AfterConnect conn hook
	AfterConnect(ctx context.Context, in *AfterConnectReq, opts ...grpc.CallOption) (*CommonResp, error)
	//AfterDisconnect conn hook
	AfterDisconnect(ctx context.Context, in *AfterDisconnectReq, opts ...grpc.CallOption) (*CommonResp, error)
	//GetUserNoticeData 获取用户通知数据
	GetUserNoticeData(ctx context.Context, in *GetUserNoticeDataReq, opts ...grpc.CallOption) (*GetUserNoticeDataResp, error)
	//AckNoticeData 确认通知数据
	AckNoticeData(ctx context.Context, in *AckNoticeDataReq, opts ...grpc.CallOption) (*AckNoticeDataResp, error)
}

type noticeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoticeServiceClient(cc grpc.ClientConnInterface) NoticeServiceClient {
	return &noticeServiceClient{cc}
}

func (c *noticeServiceClient) AfterConnect(ctx context.Context, in *AfterConnectReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/AfterConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) AfterDisconnect(ctx context.Context, in *AfterDisconnectReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/AfterDisconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) GetUserNoticeData(ctx context.Context, in *GetUserNoticeDataReq, opts ...grpc.CallOption) (*GetUserNoticeDataResp, error) {
	out := new(GetUserNoticeDataResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/GetUserNoticeData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeServiceClient) AckNoticeData(ctx context.Context, in *AckNoticeDataReq, opts ...grpc.CallOption) (*AckNoticeDataResp, error) {
	out := new(AckNoticeDataResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/AckNoticeData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoticeServiceServer is the server API for NoticeService service.
// All implementations must embed UnimplementedNoticeServiceServer
// for forward compatibility
type NoticeServiceServer interface {
	//AfterConnect conn hook
	AfterConnect(context.Context, *AfterConnectReq) (*CommonResp, error)
	//AfterDisconnect conn hook
	AfterDisconnect(context.Context, *AfterDisconnectReq) (*CommonResp, error)
	//GetUserNoticeData 获取用户通知数据
	GetUserNoticeData(context.Context, *GetUserNoticeDataReq) (*GetUserNoticeDataResp, error)
	//AckNoticeData 确认通知数据
	AckNoticeData(context.Context, *AckNoticeDataReq) (*AckNoticeDataResp, error)
	mustEmbedUnimplementedNoticeServiceServer()
}

// UnimplementedNoticeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNoticeServiceServer struct {
}

func (UnimplementedNoticeServiceServer) AfterConnect(context.Context, *AfterConnectReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AfterConnect not implemented")
}
func (UnimplementedNoticeServiceServer) AfterDisconnect(context.Context, *AfterDisconnectReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AfterDisconnect not implemented")
}
func (UnimplementedNoticeServiceServer) GetUserNoticeData(context.Context, *GetUserNoticeDataReq) (*GetUserNoticeDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserNoticeData not implemented")
}
func (UnimplementedNoticeServiceServer) AckNoticeData(context.Context, *AckNoticeDataReq) (*AckNoticeDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckNoticeData not implemented")
}
func (UnimplementedNoticeServiceServer) mustEmbedUnimplementedNoticeServiceServer() {}

// UnsafeNoticeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoticeServiceServer will
// result in compilation errors.
type UnsafeNoticeServiceServer interface {
	mustEmbedUnimplementedNoticeServiceServer()
}

func RegisterNoticeServiceServer(s grpc.ServiceRegistrar, srv NoticeServiceServer) {
	s.RegisterService(&NoticeService_ServiceDesc, srv)
}

func _NoticeService_AfterConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).AfterConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/AfterConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).AfterConnect(ctx, req.(*AfterConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_AfterDisconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AfterDisconnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).AfterDisconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/AfterDisconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).AfterDisconnect(ctx, req.(*AfterDisconnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_GetUserNoticeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserNoticeDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).GetUserNoticeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/GetUserNoticeData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).GetUserNoticeData(ctx, req.(*GetUserNoticeDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoticeService_AckNoticeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckNoticeDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).AckNoticeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/AckNoticeData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).AckNoticeData(ctx, req.(*AckNoticeDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NoticeService_ServiceDesc is the grpc.ServiceDesc for NoticeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoticeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.noticeService",
	HandlerType: (*NoticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AfterConnect",
			Handler:    _NoticeService_AfterConnect_Handler,
		},
		{
			MethodName: "AfterDisconnect",
			Handler:    _NoticeService_AfterDisconnect_Handler,
		},
		{
			MethodName: "GetUserNoticeData",
			Handler:    _NoticeService_GetUserNoticeData_Handler,
		},
		{
			MethodName: "AckNoticeData",
			Handler:    _NoticeService_AckNoticeData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notice.proto",
}
