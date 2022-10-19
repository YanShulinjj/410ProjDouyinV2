package logic

import (
	"410proj/apps/like/rpc/internal/svc"
	"410proj/apps/like/rpc/rpc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLikeLogic {
	return &IsLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsLikeLogic) IsLike(in *rpc.IsLikeReq) (*rpc.IsLikeResp, error) {
	// todo: add your logic here and delete this line
	// 查询是否包含userid 和video_id 的记录
	_, err := l.svcCtx.LikeModel.FindOneByUserIdVideoId(l.ctx, int64(in.UserId), int64(in.VideoId))
	resp := &rpc.IsLikeResp{IsLike: false}
	if err == nil {
		resp.IsLike = true
	}
	return resp, nil
}
