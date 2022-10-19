package logic

import (
	"410proj/apps/relation/rpc/model"
	"context"

	"410proj/apps/relation/rpc/internal/svc"
	"410proj/apps/relation/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *rpc.User) (*rpc.FollowList, error) {
	followlistDB, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.UserId, false)
	if err != nil {
		if err == model.ErrNotFound {
			return &rpc.FollowList{
				Follows: &rpc.Users{
					UserIds: "",
				},
			}, nil
		}
		return nil, err

	}
	return &rpc.FollowList{
		Follows: &rpc.Users{
			UserIds: followlistDB.ToUserIds,
		},
	}, nil
}
