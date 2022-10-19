package logic

import (
	"context"

	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *rpc.UpdateCommentReq) (*rpc.UpdateCommentResp, error) {
	// todo: add your logic here and delete this line

	return &rpc.UpdateCommentResp{}, nil
}
