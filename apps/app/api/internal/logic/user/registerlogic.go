package user

import (
	"410proj/apps/user/rpc/user"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"context"
	"time"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 直接调用注册用户的rpc服务
	var rpcReq user.AddUserReq
	rpcReq.UserName = req.Username
	rpcReq.Password = req.Password

	res, err := l.svcCtx.UserRPC.Register(l.ctx, &rpcReq)
	if err != nil {
		return &types.RegisterResp{
			UserResponse: types.UserResponse{
				StatusCode: int32(xerr.UserExistedErr),
				Msg:        "该用户名已经注册",
			},
		}, nil
	}
	// 注册成功
	// 生成 token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, res.UserId)
	if err != nil {
		return &types.RegisterResp{
			UserResponse: types.UserResponse{
				StatusCode: int32(xerr.TokenGenerateErr),
				Msg:        "生成Token失败",
			},
		}, nil
	}

	return &types.RegisterResp{
		UserId: res.UserId,
		Token:  accessToken,
	}, nil

}
