package logic

import (
	"410proj/apps/user/rpc/model"
	"410proj/pkg/encryption"
	"410proj/pkg/xerr"
	"context"
	"github.com/pkg/errors"

	"410proj/apps/user/rpc/internal/svc"
	"410proj/apps/user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *rpc.AddUserReq) (*rpc.AddUserResp, error) {
	// 先判断，如果用户名已经存在，返回
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.UserName)
	if err == nil {
		// 已经注册过，对比密码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserExistedErr),
			"注册已存在的用户名，username:%s,err:%v", in.UserName, err)

	} else if err != model.ErrNotFound {
		// 出现错误
		return nil, err
	} else {
		// 注册新用户
		md5str, err := encryption.Md5ByString(in.Password)
		if err != nil {
			return nil, err
		}
		user := model.User{
			Username: in.UserName,
			Password: md5str,
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
		if err != nil {
			return nil, err
		}
		var resp rpc.AddUserResp
		userId, _ := res.LastInsertId()
		resp.UserId = uint64(userId)
		return &resp, nil
	}

	return &rpc.AddUserResp{}, nil
}
