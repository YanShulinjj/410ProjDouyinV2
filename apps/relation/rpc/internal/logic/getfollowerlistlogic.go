package logic

import (
	"410proj/apps/relation/rpc/model"
	"context"

	"410proj/apps/relation/rpc/internal/svc"
	"410proj/apps/relation/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *rpc.User) (*rpc.FollowerList, error) {

	followerlistDB, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.UserId, true)
	if err != nil {
		if err == model.ErrNotFound {
			return &rpc.FollowerList{
				Followers: &rpc.Users{
					UserIds: "",
				},
			}, nil
		}
		return nil, err
	}
	return &rpc.FollowerList{
		Followers: &rpc.Users{
			UserIds: followerlistDB.ToUserIds,
		},
	}, nil
}
