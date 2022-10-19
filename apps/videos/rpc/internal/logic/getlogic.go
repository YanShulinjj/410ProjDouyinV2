package logic

import (
	"context"

	"410proj/apps/videos/rpc/internal/svc"
	"410proj/apps/videos/rpc/rpc"

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

func (l *GetLogic) Get(in *rpc.VideoItemReq) (*rpc.VideoItem, error) {
	videoDB, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	return &rpc.VideoItem{
		Id:         videoDB.VideoId,
		AuthorId:   videoDB.UserId,
		PlayUrl:    videoDB.PlayUrl,
		CoverUrl:   videoDB.CoverUrl,
		CreateTime: videoDB.CreateTime.UnixNano(),
	}, nil
}
