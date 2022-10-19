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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 对登陆用户的rpc服务进行调用
	var rpcReq user.UserLoginReq
	rpcReq.UserName = req.Username
	rpcReq.Password = req.Password

	res, err := l.svcCtx.UserRPC.Login(l.ctx, &rpcReq)
	if err != nil {
		return &types.LoginResp{
			UserResponse: types.UserResponse{
				StatusCode: int32(xerr.UserPwdNotMatchErr),
				Msg:        "用户名或密码错误",
			},
		}, nil
	}
	// 登陆成功
	// 生成 token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, res.UserId)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		UserId: res.UserId,
		Token:  accessToken,
	}, nil

}
