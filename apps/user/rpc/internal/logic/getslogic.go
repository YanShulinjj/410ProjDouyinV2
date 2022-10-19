package logic

import (
	"context"
	"strconv"
	"strings"

	"410proj/apps/user/rpc/internal/svc"
	"410proj/apps/user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetsLogic {
	return &GetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetsLogic) Gets(in *rpc.UserReq) (*rpc.UserResp, error) {
	// 通过多个uid 获取用户列表
	uidstrs := strings.Split(in.UserIds, ",")
	users := map[uint64]*rpc.UserItem{}
	for _, uidstr := range uidstrs {
		uid, err := strconv.Atoi(uidstr)
		if err != nil {
			return nil, err
		}
		userDB, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(uid))
		if err != nil {
			return nil, err
		}
		useritem := rpc.UserItem{
			UserId:   uint64(userDB.UserId),
			UserName: userDB.Username,
			// Password: userDB.Password,
		}
		users[useritem.UserId] = &useritem
	}
	return &rpc.UserResp{
		Users: users,
	}, nil
}
