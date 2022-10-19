// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package server

import (
	"context"

	"410proj/apps/comment/rpc/internal/logic"
	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) Get(ctx context.Context, in *rpc.CommentItemReq) (*rpc.CommentItem, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *CommentServer) GetCommentNum(ctx context.Context, in *rpc.CommentNumReq) (*rpc.CommentNumResp, error) {
	l := logic.NewGetCommentNumLogic(ctx, s.svcCtx)
	return l.GetCommentNum(in)
}

func (s *CommentServer) Gets(ctx context.Context, in *rpc.CommentReq) (*rpc.CommentResp, error) {
	l := logic.NewGetsLogic(ctx, s.svcCtx)
	return l.Gets(in)
}

func (s *CommentServer) GetByPage(ctx context.Context, in *rpc.CommentPageReq) (*rpc.CommentPageResp, error) {
	l := logic.NewGetByPageLogic(ctx, s.svcCtx)
	return l.GetByPage(in)
}

func (s *CommentServer) Add(ctx context.Context, in *rpc.AddCommentReq) (*rpc.AddCommentResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *CommentServer) Update(ctx context.Context, in *rpc.UpdateCommentReq) (*rpc.UpdateCommentResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *CommentServer) Drop(ctx context.Context, in *rpc.DropCommentReq) (*rpc.DropCommentResp, error) {
	l := logic.NewDropLogic(ctx, s.svcCtx)
	return l.Drop(in)
}
