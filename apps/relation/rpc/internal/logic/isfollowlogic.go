package logic

import (
	"410proj/apps/relation/rpc/model"
	"context"
	"strconv"
	"strings"

	"410proj/apps/relation/rpc/internal/svc"
	"410proj/apps/relation/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowLogic {
	return &IsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFollowLogic) IsFollow(in *rpc.IsFollowReq) (*rpc.IsFollowResp, error) {
	//
	userIds, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.UserId, false)
	if err != nil {
		if err == model.ErrNotFound {
			return &rpc.IsFollowResp{
				IsFollow: false,
			}, nil
		}
		return nil, err
	}
	useridstrs := strings.Split(userIds.ToUserIds, ",")
	for _, useridstr := range useridstrs {
		userid, _ := strconv.Atoi(useridstr)
		if userid == int(in.ToUserId) {
			return &rpc.IsFollowResp{
				IsFollow: true,
			}, nil
		}
	}
	return &rpc.IsFollowResp{}, nil
}
