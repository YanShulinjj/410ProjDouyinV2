package logic

import (
	"context"

	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentNumLogic {
	return &GetCommentNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentNumLogic) GetCommentNum(in *rpc.CommentNumReq) (*rpc.CommentNumResp, error) {

	res, err := l.svcCtx.CommentNumModel.FindOneByVideoId(l.ctx, int64(in.VideoId))
	if err != nil {
		return nil, err
	}

	return &rpc.CommentNumResp{
		Nums: uint64(res.CommentNum),
	}, nil
}
