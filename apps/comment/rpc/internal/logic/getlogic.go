package logic

import (
	"context"

	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *rpc.CommentItemReq) (*rpc.CommentItem, error) {
	commentDB, err := l.svcCtx.CommentModel.FindOne(l.ctx, int64(in.CommentId))
	if err != nil {
		return nil, err
	}
	return &rpc.CommentItem{
		CommentId:  uint64(commentDB.CommentId),
		UserId:     uint64(commentDB.UserId),
		VideoId:    uint64(commentDB.VideoId),
		Content:    commentDB.Content,
		CreateTime: commentDB.CreateTime.UnixNano(),
	}, nil
}
