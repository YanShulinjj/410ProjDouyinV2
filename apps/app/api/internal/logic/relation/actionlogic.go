package relation

import (
	"410proj/apps/relation/rpc/relation"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"context"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLogic) Action(req *types.ActionRelationReq) (resp *types.ActionRelationResp, err error) {
	// action_type = 1 关注
	//             = 2 取消关注
	// 通过token获取userid
	userId, err := jwtx.GetUserId(l.svcCtx.Config.JwtAuth.AccessSecret, req.Token)
	if err != nil {
		return &types.ActionRelationResp{
			LikeResponse: types.LikeResponse{
				StatusCode: int32(xerr.ReuqestParamErr),
				Msg:        "获取UserId失败",
			}}, nil
	}

	if req.ActionType == 1 {
		relationRpcReq := relation.FollowReq{
			UserId:   userId,
			ToUserId: req.ToUserId,
		}
		_, err := l.svcCtx.RelationRPC.Follow(l.ctx, &relationRpcReq)
		if err != nil {
			return &types.ActionRelationResp{
				LikeResponse: types.LikeResponse{
					StatusCode: int32(xerr.ReuqestParamErr),
					Msg:        "关注失败",
				}}, nil
		}
	} else if req.ActionType == 2 {
		// 取消关注
		relationRpcReq := relation.CancelFollowReq{
			UserId:   userId,
			ToUserId: req.ToUserId,
		}
		_, err := l.svcCtx.RelationRPC.CancelFollow(l.ctx, &relationRpcReq)
		if err != nil {
			return &types.ActionRelationResp{
				LikeResponse: types.LikeResponse{
					StatusCode: int32(xerr.ReuqestParamErr),
					Msg:        "取消关注失败",
				}}, nil
		}
	}
	return &types.ActionRelationResp{
		LikeResponse: types.LikeResponse{
			StatusCode: 0,
			Msg:        "ok",
		}}, nil
}
