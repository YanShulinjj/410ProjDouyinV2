package relation

import (
	"410proj/apps/app/api/internal/logic/user"
	"410proj/apps/relation/rpc/relation"
	"410proj/pkg/xerr"
	"context"
	"strconv"
	"strings"
	"sync"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowlistLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	userApi *user.UserinfoLogic
}

func NewFollowlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowlistLogic {
	return &FollowlistLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		userApi: user.NewUserinfoLogic(ctx, svcCtx),
	}
}

func (l *FollowlistLogic) Followlist(req *types.FollowListReq) (resp *types.FollowListResp, err error) {
	relationRpcReq := relation.User{
		UserId: req.UserId,
	}
	res, err := l.svcCtx.RelationRPC.GetFollowList(l.ctx, &relationRpcReq)
	if err != nil {
		return &types.FollowListResp{
			RelationResponse: types.RelationResponse{
				StatusCode: int32(xerr.ReuqestParamErr),
				Msg:        "获取关注列表失败",
			}}, nil
	}

	uidstrs := strings.Split(res.Follows.UserIds, ",")

	if len(uidstrs) == 0 {
		return &types.FollowListResp{
			RelationResponse: types.RelationResponse{
				StatusCode: 0,
				Msg:        "ok",
			},
		}, nil
	}
	users := make([]*types.UserInfo, len(uidstrs))

	// 此处采用多线程获取信息
	wg := sync.WaitGroup{}
	for i, uidstr := range uidstrs {
		wg.Add(1)
		go func(j int, uid string) {
			defer wg.Done()
			id, err := strconv.Atoi(uid)
			if err != nil {
				return
			}
			r, err := l.userApi.Userinfo(&types.UserInfoReq{
				UserId: uint64(id),
			})
			if err != nil {
				return
			}
			// 通过 rpc 获取 IsFollow 字段
			isfollow := false
			isfollowResp, err := l.svcCtx.RelationRPC.IsFollow(l.ctx, &relation.IsFollowReq{
				UserId:   req.UserId,
				ToUserId: r.UserInfo.Id,
			})
			if err == nil {
				isfollow = isfollowResp.IsFollow
			}
			// 通过rpc 获取followcount 字段
			var followcount, followercount int64
			followcountResp, err := l.svcCtx.RelationRPC.FollowCount(l.ctx, &relation.User{
				UserId: req.UserId,
			})
			if err == nil {
				followcount = int64(followcountResp.Count)
			}

			followercountResp, err := l.svcCtx.RelationRPC.FollowerCount(l.ctx, &relation.User{
				UserId: req.UserId,
			})
			if err == nil {
				followercount = int64(followercountResp.Count)
			}

			userItem := &types.UserInfo{
				Id:            r.UserInfo.Id,
				Username:      r.UserInfo.Username,
				IsFollow:      isfollow,
				FollowerCount: followcount,
				FollowCount:   followercount,
			}
			users[j] = userItem
		}(i, uidstr)
	}
	wg.Wait()

	return &types.FollowListResp{
		RelationResponse: types.RelationResponse{
			StatusCode: 0,
			Msg:        "ok",
		},
		Users: users,
	}, nil
}
