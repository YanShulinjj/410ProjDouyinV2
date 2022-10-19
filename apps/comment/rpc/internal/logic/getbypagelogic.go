package logic

import (
	"context"

	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByPageLogic {
	return &GetByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByPageLogic) GetByPage(in *rpc.CommentPageReq) (*rpc.CommentPageResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.CommentPageResp{}, nil
}
