// Code generated by goctl. DO NOT EDIT!
// Source: like.proto

package server

import (
	"context"

	"410proj/apps/like/rpc/internal/logic"
	"410proj/apps/like/rpc/internal/svc"
	"410proj/apps/like/rpc/rpc"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

func (s *LikeServer) GetLikeList(ctx context.Context, in *rpc.User) (*rpc.LikeList, error) {
	l := logic.NewGetLikeListLogic(ctx, s.svcCtx)
	return l.GetLikeList(in)
}

func (s *LikeServer) GetLikeNum(ctx context.Context, in *rpc.LikeNumReq) (*rpc.LikeNumResp, error) {
	l := logic.NewGetLikeNumLogic(ctx, s.svcCtx)
	return l.GetLikeNum(in)
}

func (s *LikeServer) IsLike(ctx context.Context, in *rpc.IsLikeReq) (*rpc.IsLikeResp, error) {
	l := logic.NewIsLikeLogic(ctx, s.svcCtx)
	return l.IsLike(in)
}

func (s *LikeServer) LikeVideo(ctx context.Context, in *rpc.LikeVideoReq) (*rpc.LikeVideoResp, error) {
	l := logic.NewLikeVideoLogic(ctx, s.svcCtx)
	return l.LikeVideo(in)
}

func (s *LikeServer) CancelLikeVideo(ctx context.Context, in *rpc.CancelLikeVideoReq) (*rpc.CancelLikeVideoResp, error) {
	l := logic.NewCancelLikeVideoLogic(ctx, s.svcCtx)
	return l.CancelLikeVideo(in)
}