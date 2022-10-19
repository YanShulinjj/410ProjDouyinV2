package logic

import (
	"context"

	"410proj/apps/like/rpc/internal/svc"
	"410proj/apps/like/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikeNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikeNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikeNumLogic {
	return &GetLikeNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikeNumLogic) GetLikeNum(in *rpc.LikeNumReq) (*rpc.LikeNumResp, error) {

	resDB, err := l.svcCtx.LikeNumModel.FindOneByVideoId(l.ctx, int64(in.VideoId))
	if err != nil {
		return nil, err
	}

	return &rpc.LikeNumResp{
		Nums: uint64(resDB.Likes),
	}, nil
}
