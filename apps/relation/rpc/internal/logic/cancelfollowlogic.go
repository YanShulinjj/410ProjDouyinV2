package logic

import (
	"context"
	"strconv"
	"strings"

	"410proj/apps/relation/rpc/internal/svc"
	"410proj/apps/relation/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelFollowLogic {
	return &CancelFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelFollowLogic) CancelFollow(in *rpc.CancelFollowReq) (*rpc.CancelFollowResp, error) {
	// todo: add your logic here and delete this line
	urDB, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.UserId, false)
	if err != nil {
		return nil, err
	}
	tuDB, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.ToUserId, true)
	if err != nil {
		return nil, err
	}
	// 删除ur的关注列表
	tidstrs := strings.Split(urDB.ToUserIds, ",")
	newtidstrs := make([]string, 0, len(tidstrs))
	for _, tidstr := range tidstrs {
		if strconv.Itoa(int(in.ToUserId)) != tidstr {
			newtidstrs = append(newtidstrs, tidstr)
		}
	}
	urDB.ToUserIds = strings.Join(newtidstrs, ",")
	// 删除tu的粉丝列表
	tidstrs = strings.Split(tuDB.ToUserIds, ",")
	newtidstrs = make([]string, 0, len(tidstrs))
	for _, tidstr := range tidstrs {
		if strconv.Itoa(int(in.UserId)) != tidstr {
			newtidstrs = append(newtidstrs, tidstr)
		}
	}
	tuDB.ToUserIds = strings.Join(newtidstrs, ",")

	err = l.svcCtx.RelationModel.Update(l.ctx, urDB)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RelationModel.Update(l.ctx, tuDB)
	if err != nil {
		return nil, err
	}
	return &rpc.CancelFollowResp{}, nil
}
