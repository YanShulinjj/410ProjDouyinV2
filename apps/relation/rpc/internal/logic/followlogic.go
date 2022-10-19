package logic

import (
	"410proj/apps/relation/rpc/model"
	"context"
	"strconv"

	"410proj/apps/relation/rpc/internal/svc"
	"410proj/apps/relation/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *rpc.FollowReq) (*rpc.FollowResp, error) {
	// api层需要保证未关注才进行关注
	urDB, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.UserId, false)
	if err != nil {
		if err == model.ErrNotFound {
			// 插入条目
			userRelationDB := model.Relation{
				UserId:    in.UserId,
				ToUserIds: strconv.Itoa(int(in.ToUserId)),
				Type:      false,
			}
			toUserRelationDB := model.Relation{
				UserId:    in.ToUserId,
				ToUserIds: strconv.Itoa(int(in.UserId)),
				Type:      true,
			}
			_, err = l.svcCtx.RelationModel.Insert(l.ctx, &userRelationDB)
			if err != nil {
				return nil, err
			}
			_, err = l.svcCtx.RelationModel.Insert(l.ctx, &toUserRelationDB)
			if err != nil {
				return nil, err
			}
			return &rpc.FollowResp{}, nil
		}
		return nil, err
	}
	tuDB, err := l.svcCtx.RelationModel.FindOneByUserIdType(l.ctx, in.ToUserId, true)
	if err != nil {
		if err == model.ErrNotFound {
			// 插入条目
			toUserRelationDB := model.Relation{
				UserId:    in.ToUserId,
				ToUserIds: strconv.Itoa(int(in.UserId)),
				Type:      true,
			}
			// 添加ur的关注列表
			tids := strconv.Itoa(int(in.ToUserId))
			if len(urDB.ToUserIds) > 0 {
				tids = urDB.ToUserIds + "," + tids
			}
			urDB.ToUserIds = tids
			err = l.svcCtx.RelationModel.Update(l.ctx, urDB)
			if err != nil {
				return nil, err
			}
			// 插入tr的粉丝列表
			_, err = l.svcCtx.RelationModel.Insert(l.ctx, &toUserRelationDB)
			if err != nil {
				return nil, err
			}
			return &rpc.FollowResp{}, nil
		}
		return nil, err
	}
	// 添加ur的关注列表
	tids := strconv.Itoa(int(in.ToUserId))
	if len(urDB.ToUserIds) > 0 {
		tids = urDB.ToUserIds + "," + tids
	}
	urDB.ToUserIds = tids
	// 添加tu的粉丝列表
	tids = strconv.Itoa(int(in.UserId))
	if len(tuDB.ToUserIds) > 0 {
		tids = tuDB.ToUserIds + "," + tids
	}
	tuDB.ToUserIds = tids

	err = l.svcCtx.RelationModel.Update(l.ctx, urDB)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RelationModel.Update(l.ctx, tuDB)
	if err != nil {
		return nil, err
	}

	return &rpc.FollowResp{}, nil
}
