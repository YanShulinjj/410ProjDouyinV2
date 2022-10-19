package logic

import (
	"410proj/pkg/encryption"
	"410proj/pkg/xerr"
	"context"
	"github.com/pkg/errors"

	"410proj/apps/user/rpc/internal/svc"
	"410proj/apps/user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *rpc.UpdateUserReq) (*rpc.UpdateUserResp, error) {
	// 判断newpassword字段是否为空字符串
	if len(in.Newpassword) == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ReuqestParamErr),
			"修改密码发生错误。")
	}
	// 修改用户密码
	// 1.首先查询是否存在该用户
	userDB, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.UserName)
	if err != nil {
		return nil, err
	}
	// 2.判断旧密码是否正确
	old := in.Oldpassword
	if oldmd5, _ := encryption.Md5ByString(old); oldmd5 != userDB.Password {
		// 旧密码错误
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserPwdNotMatchErr),
			"修改密码发生错误。")
	}
	newpwd, err := encryption.Md5ByString(in.Newpassword)
	if err != nil {
		return nil, err
	}
	userDB.Password = newpwd
	err = l.svcCtx.UserModel.Update(l.ctx, userDB)
	if err != nil {
		return nil, err
	}
	return &rpc.UpdateUserResp{}, nil
}
