package logic

import (
	"410proj/apps/user/rpc/internal/svc"
	"410proj/apps/user/rpc/model"
	"410proj/apps/user/rpc/rpc"
	"410proj/apps/user/rpc/user"
	"410proj/pkg/encryption"
	"410proj/pkg/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *rpc.UserLoginReq) (*rpc.UserLoginResp, error) {
	// 先判断用户是否存在
	userDB, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotExistErr),
				"根据username查询用户信息失败，username:%s,err:%v", in.UserName, err)
		}
		return nil, err
	}
	// 验证密码
	md5str, err := encryption.Md5ByString(in.Password)
	if md5str != userDB.Password {
		return nil, errors.Wrapf(xerr.NewErrMsg("账号或密码错误"),
			"密码错误")
	}
	// 成功登陆
	var resp user.UserLoginResp
	_ = copier.Copy(&resp, userDB)

	return &resp, nil
}
