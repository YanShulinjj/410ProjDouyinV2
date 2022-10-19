package logic

import (
	"410proj/apps/like/rpc/model"
	"context"

	"410proj/apps/like/rpc/internal/svc"
	"410proj/apps/like/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeVideoLogic) LikeVideo(in *rpc.LikeVideoReq) (*rpc.LikeVideoResp, error) {
	// 对视频点赞
	// 上层需要保证: video_id 和 user_id 是合法的
	// 后续：需要事务支持？
	// 1. 向点赞映射表添加条目
	userlike := model.UserLikeMap{
		UserId:  int64(in.UserId),
		VideoId: int64(in.VideoId),
	}
	_, err := l.svcCtx.LikeModel.Insert(l.ctx, &userlike)
	if err != nil {
		return nil, err
	}
	// 2. 先获取点赞数目表
	likenumDB, err := l.svcCtx.LikeNumModel.FindOneByVideoId(l.ctx, int64(in.VideoId))
	if err != nil {
		if err == model.ErrNotFound {
			// 新建条目
			likenum := model.VideoLikeNum{
				VideoId: int64(in.VideoId),
				Likes:   1,
			}
			_, err = l.svcCtx.LikeNumModel.Insert(l.ctx, &likenum)
			if err != nil {
				return nil, err
			}
			return &rpc.LikeVideoResp{}, nil
		}
		return nil, err
	}

	likenumDB.Likes++
	// 再更新到DB
	err = l.svcCtx.LikeNumModel.Update(l.ctx, likenumDB)
	if err != nil {
		return nil, err
	}

	return &rpc.LikeVideoResp{}, nil
}
