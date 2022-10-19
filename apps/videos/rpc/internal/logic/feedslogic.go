package logic

import (
	"context"
	"time"

	"410proj/apps/videos/rpc/internal/svc"
	"410proj/apps/videos/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

const VideoFeedNum = 10

type FeedsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedsLogic {
	return &FeedsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedsLogic) Feeds(in *rpc.VideoFeedReq) (*rpc.VideoFeedResp, error) {
	// 获取在请求时间之前的前N个视频
	timestr := time.Now().Format("2006-01-02 15:04:05")
	videosDB, err := l.svcCtx.VideoModel.FindManyBefore(l.ctx, timestr, VideoFeedNum)
	if err != nil {
		return nil, err
	}
	videos := []*rpc.VideoItem{}
	for _, videoDB := range videosDB {
		video := &rpc.VideoItem{
			Id:       videoDB.VideoId,
			AuthorId: videoDB.UserId,
			PlayUrl:  videoDB.PlayUrl,
			CoverUrl: videoDB.CoverUrl,
		}
		videos = append(videos, video)
	}
	return &rpc.VideoFeedResp{
		Videos: videos,
	}, nil
}
