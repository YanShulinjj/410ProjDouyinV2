// Code generated by goctl. DO NOT EDIT!
// Source: video.proto

package server

import (
	"context"

	"410proj/apps/videos/rpc/internal/logic"
	"410proj/apps/videos/rpc/internal/svc"
	"410proj/apps/videos/rpc/rpc"
)

type VideoServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedVideoServer
}

func NewVideoServer(svcCtx *svc.ServiceContext) *VideoServer {
	return &VideoServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServer) Get(ctx context.Context, in *rpc.VideoItemReq) (*rpc.VideoItem, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *VideoServer) Gets(ctx context.Context, in *rpc.VideoReq) (*rpc.VideoResp, error) {
	l := logic.NewGetsLogic(ctx, s.svcCtx)
	return l.Gets(in)
}

func (s *VideoServer) PublishList(ctx context.Context, in *rpc.VideoPublishReq) (*rpc.VideoPublishResp, error) {
	l := logic.NewPublishListLogic(ctx, s.svcCtx)
	return l.PublishList(in)
}

func (s *VideoServer) Feeds(ctx context.Context, in *rpc.VideoFeedReq) (*rpc.VideoFeedResp, error) {
	l := logic.NewFeedsLogic(ctx, s.svcCtx)
	return l.Feeds(in)
}

func (s *VideoServer) Add(ctx context.Context, in *rpc.AddVideoReq) (*rpc.AddVideoResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}
