package favorite

import (
	"410proj/apps/like/rpc/like"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"context"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.ActionLikeReq) (resp *types.ActionLikeResp, err error) {
	//
	// action_type = 1 点赞
	// 通过token获取userid
	userId, err := jwtx.GetUserId(l.svcCtx.Config.JwtAuth.AccessSecret, req.Token)
	if err != nil {
		return &types.ActionLikeResp{
			LikeResponse: types.LikeResponse{
				StatusCode: int32(xerr.ReuqestParamErr),
				Msg:        "点赞失败",
			}}, nil
	}

	if req.ActionType == 1 {
		likeRpcReq := like.LikeVideoReq{
			UserId:  userId,
			VideoId: req.VideoId,
		}
		_, err := l.svcCtx.LikeRPC.LikeVideo(l.ctx, &likeRpcReq)
		if err != nil {
			return &types.ActionLikeResp{
				LikeResponse: types.LikeResponse{
					StatusCode: int32(xerr.ReuqestParamErr),
					Msg:        "点赞失败",
				}}, nil
		}
	} else if req.ActionType == 2 {
		// 取消点赞
		likeRpcReq := like.CancelLikeVideoReq{
			UserId:  userId,
			VideoId: req.VideoId,
		}
		_, err := l.svcCtx.LikeRPC.CancelLikeVideo(l.ctx, &likeRpcReq)
		if err != nil {
			return &types.ActionLikeResp{
				LikeResponse: types.LikeResponse{
					StatusCode: int32(xerr.ReuqestParamErr),
					Msg:        "取消点赞失败",
				}}, nil
		}
	}

	return &types.ActionLikeResp{
		LikeResponse: types.LikeResponse{
			StatusCode: 0,
			Msg:        "ok",
		}}, nil
}
