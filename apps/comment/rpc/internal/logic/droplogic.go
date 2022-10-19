package logic

import (
	"context"

	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropLogic {
	return &DropLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DropLogic) Drop(in *rpc.DropCommentReq) (*rpc.DropCommentResp, error) {
	// 删除一条评论，请在api层做验证

	err := l.svcCtx.CommentModel.Delete(l.ctx, int64(in.CommentId), int64(in.VideoId))
	if err != nil {
		return nil, err
	}
	commentNumDB, err := l.svcCtx.CommentNumModel.FindOneByVideoId(l.ctx, int64(in.VideoId))
	if err != nil {
		return nil, err
	}
	commentNumDB.CommentNum--
	if err != nil {
		return nil, err
	}
	// 再更新到DB
	err = l.svcCtx.CommentNumModel.Update(l.ctx, commentNumDB)
	if err != nil {
		return nil, err
	}
	return &rpc.DropCommentResp{}, nil
}
