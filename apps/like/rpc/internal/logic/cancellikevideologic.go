package logic

import (
	"context"

	"410proj/apps/like/rpc/internal/svc"
	"410proj/apps/like/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeVideoLogic {
	return &CancelLikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLikeVideoLogic) CancelLikeVideo(in *rpc.CancelLikeVideoReq) (*rpc.CancelLikeVideoResp, error) {
	// 对视频取消点赞
	// 上层需要保证: video_id 和 user_id 是合法的
	// 后续：需要事务支持？
	// 1. 向点赞映射表删除条目
	userlikeDB, err := l.svcCtx.LikeModel.FindOneByUserIdVideoId(l.ctx, int64(in.UserId), int64(in.VideoId))
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.LikeModel.Delete(l.ctx, userlikeDB.LikeId)
	if err != nil {
		return nil, err
	}
	// 2. 先获取点赞数目表
	likenumDB, err := l.svcCtx.LikeNumModel.FindOneByVideoId(l.ctx, int64(in.VideoId))
	if err != nil {
		return nil, err
	}
	likenumDB.Likes--
	// 3. 再更新到DB
	err = l.svcCtx.LikeNumModel.Update(l.ctx, likenumDB)
	if err != nil {
		return nil, err
	}

	return &rpc.CancelLikeVideoResp{}, nil
}
