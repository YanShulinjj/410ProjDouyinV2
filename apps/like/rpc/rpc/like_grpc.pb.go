// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: like.proto

package rpc

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

// LikeClient is the client API for Like service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeClient interface {
	GetLikeList(ctx context.Context, in *User, opts ...grpc.CallOption) (*LikeList, error)
	GetLikeNum(ctx context.Context, in *LikeNumReq, opts ...grpc.CallOption) (*LikeNumResp, error)
	IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeResp, error)
	LikeVideo(ctx context.Context, in *LikeVideoReq, opts ...grpc.CallOption) (*LikeVideoResp, error)
	CancelLikeVideo(ctx context.Context, in *CancelLikeVideoReq, opts ...grpc.CallOption) (*CancelLikeVideoResp, error)
}

type likeClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeClient(cc grpc.ClientConnInterface) LikeClient {
	return &likeClient{cc}
}

func (c *likeClient) GetLikeList(ctx context.Context, in *User, opts ...grpc.CallOption) (*LikeList, error) {
	out := new(LikeList)
	err := c.cc.Invoke(ctx, "/like.Like/GetLikeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) GetLikeNum(ctx context.Context, in *LikeNumReq, opts ...grpc.CallOption) (*LikeNumResp, error) {
	out := new(LikeNumResp)
	err := c.cc.Invoke(ctx, "/like.Like/GetLikeNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeResp, error) {
	out := new(IsLikeResp)
	err := c.cc.Invoke(ctx, "/like.Like/IsLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) LikeVideo(ctx context.Context, in *LikeVideoReq, opts ...grpc.CallOption) (*LikeVideoResp, error) {
	out := new(LikeVideoResp)
	err := c.cc.Invoke(ctx, "/like.Like/LikeVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) CancelLikeVideo(ctx context.Context, in *CancelLikeVideoReq, opts ...grpc.CallOption) (*CancelLikeVideoResp, error) {
	out := new(CancelLikeVideoResp)
	err := c.cc.Invoke(ctx, "/like.Like/CancelLikeVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServer is the server API for Like service.
// All implementations must embed UnimplementedLikeServer
// for forward compatibility
type LikeServer interface {
	GetLikeList(context.Context, *User) (*LikeList, error)
	GetLikeNum(context.Context, *LikeNumReq) (*LikeNumResp, error)
	IsLike(context.Context, *IsLikeReq) (*IsLikeResp, error)
	LikeVideo(context.Context, *LikeVideoReq) (*LikeVideoResp, error)
	CancelLikeVideo(context.Context, *CancelLikeVideoReq) (*CancelLikeVideoResp, error)
	mustEmbedUnimplementedLikeServer()
}

// UnimplementedLikeServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServer struct {
}

func (UnimplementedLikeServer) GetLikeList(context.Context, *User) (*LikeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikeList not implemented")
}
func (UnimplementedLikeServer) GetLikeNum(context.Context, *LikeNumReq) (*LikeNumResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikeNum not implemented")
}
func (UnimplementedLikeServer) IsLike(context.Context, *IsLikeReq) (*IsLikeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLike not implemented")
}
func (UnimplementedLikeServer) LikeVideo(context.Context, *LikeVideoReq) (*LikeVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeVideo not implemented")
}
func (UnimplementedLikeServer) CancelLikeVideo(context.Context, *CancelLikeVideoReq) (*CancelLikeVideoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLikeVideo not implemented")
}
func (UnimplementedLikeServer) mustEmbedUnimplementedLikeServer() {}

// UnsafeLikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServer will
// result in compilation errors.
type UnsafeLikeServer interface {
	mustEmbedUnimplementedLikeServer()
}

func RegisterLikeServer(s grpc.ServiceRegistrar, srv LikeServer) {
	s.RegisterService(&Like_ServiceDesc, srv)
}

func _Like_GetLikeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetLikeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.Like/GetLikeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetLikeList(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_GetLikeNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetLikeNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.Like/GetLikeNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetLikeNum(ctx, req.(*LikeNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_IsLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).IsLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.Like/IsLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).IsLike(ctx, req.(*IsLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_LikeVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).LikeVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.Like/LikeVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).LikeVideo(ctx, req.(*LikeVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_CancelLikeVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLikeVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).CancelLikeVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/like.Like/CancelLikeVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).CancelLikeVideo(ctx, req.(*CancelLikeVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Like_ServiceDesc is the grpc.ServiceDesc for Like service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Like_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "like.Like",
	HandlerType: (*LikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLikeList",
			Handler:    _Like_GetLikeList_Handler,
		},
		{
			MethodName: "GetLikeNum",
			Handler:    _Like_GetLikeNum_Handler,
		},
		{
			MethodName: "IsLike",
			Handler:    _Like_IsLike_Handler,
		},
		{
			MethodName: "LikeVideo",
			Handler:    _Like_LikeVideo_Handler,
		},
		{
			MethodName: "CancelLikeVideo",
			Handler:    _Like_CancelLikeVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like.proto",
}
