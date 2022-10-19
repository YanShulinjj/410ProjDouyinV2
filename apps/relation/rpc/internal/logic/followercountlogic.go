package logic

import (
	"410proj/apps/relation/rpc/model"
	"context"
	"strings"

	"410proj/apps/relation/rpc/internal/svc"
	"410proj/apps/relation/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerCountLogic {
	return &FollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerCountLogic) FollowerCount(in *rpc.User) (*rpc.UserCount, error) {
	userIds, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.UserId, true)
	if err != nil {
		if err == model.ErrNotFound {
			return &rpc.UserCount{
				Count: 0,
			}, nil
		}
		return nil, err
	}
	useridstrs := strings.Split(userIds.ToUserIds, ",")

	return &rpc.UserCount{
		Count: uint64(len(useridstrs)),
	}, nil
}
