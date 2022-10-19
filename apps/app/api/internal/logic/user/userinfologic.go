package user

import (
	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"
	"410proj/apps/user/rpc/user"
	"410proj/pkg/xerr"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	//
	var rpcReq user.UserItemReq
	rpcReq.UserId = req.UserId

	res, err := l.svcCtx.UserRPC.Get(l.ctx, &rpcReq)
	if err != nil {
		return &types.UserInfoResp{
			UserResponse: types.UserResponse{
				StatusCode: int32(xerr.UserNotExistErr),
				Msg:        "user doesn't exits!",
			},
		}, nil
	}
	// 注意此处需要添加获得isfollow的字段值
	return &types.UserInfoResp{
		UserInfo: types.UserInfo{
			Id:       res.UserId,
			Username: res.UserName,
			IsFollow: false,
		},
	}, nil
}
