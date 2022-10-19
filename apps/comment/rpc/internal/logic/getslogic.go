package logic

import (
	"410proj/apps/comment/rpc/comment"
	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetsLogic {
	return &GetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetsLogic) Gets(in *rpc.CommentReq) (*rpc.CommentResp, error) {

	commentsDB, err := l.svcCtx.CommentModel.FindMany(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	comments := make([]*comment.CommentItem, len(commentsDB))
	for i, commentDB := range commentsDB {
		commentitem := &rpc.CommentItem{
			CommentId:  uint64(commentDB.CommentId),
			UserId:     uint64(commentDB.UserId),
			VideoId:    uint64(commentDB.VideoId),
			Content:    commentDB.Content,
			CreateTime: commentDB.CreateTime.Unix(),
		}
		comments[len(comments)-1-i] = commentitem
	}

	return &rpc.CommentResp{
		Comments: comments,
	}, nil
}
