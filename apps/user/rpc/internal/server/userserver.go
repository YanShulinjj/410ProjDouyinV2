// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"410proj/apps/user/rpc/internal/logic"
	"410proj/apps/user/rpc/internal/svc"
	"410proj/apps/user/rpc/rpc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Get(ctx context.Context, in *rpc.UserItemReq) (*rpc.UserItem, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *UserServer) Gets(ctx context.Context, in *rpc.UserReq) (*rpc.UserResp, error) {
	l := logic.NewGetsLogic(ctx, s.svcCtx)
	return l.Gets(in)
}

func (s *UserServer) Register(ctx context.Context, in *rpc.AddUserReq) (*rpc.AddUserResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *rpc.UserLoginReq) (*rpc.UserLoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Update(ctx context.Context, in *rpc.UpdateUserReq) (*rpc.UpdateUserResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}
