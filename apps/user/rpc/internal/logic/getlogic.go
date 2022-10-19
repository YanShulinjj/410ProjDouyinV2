package logic

import (
	"410proj/apps/user/rpc/model"
	"410proj/pkg/xerr"
	"context"
	"github.com/pkg/errors"

	"410proj/apps/user/rpc/internal/svc"
	"410proj/apps/user/rpc/rpc"

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

func (l *GetLogic) Get(in *rpc.UserItemReq) (*rpc.UserItem, error) {
	// 获取用户信息
	userDB, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(in.UserId))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotExistErr),
				"根据username查询用户信息失败，user_id:%s,err:%v", in.UserId, err)
		}
		return nil, err
	}
	return &rpc.UserItem{
		UserId:   uint64(userDB.UserId),
		UserName: userDB.Username,
	}, nil
}
