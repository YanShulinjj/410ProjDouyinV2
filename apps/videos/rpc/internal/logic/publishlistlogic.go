package logic

import (
	"410proj/apps/videos/rpc/video"
	"context"
	"strconv"
	"strings"

	"410proj/apps/videos/rpc/internal/svc"
	"410proj/apps/videos/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	gets *GetsLogic
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		gets:   NewGetsLogic(ctx, svcCtx),
	}
}

func (l *PublishListLogic) PublishList(in *rpc.VideoPublishReq) (*rpc.VideoPublishResp, error) {
	//
	videoIdsDB, err := l.svcCtx.PublishModel.FindMany(l.ctx, int64(in.UserId))
	if err != nil {
		return nil, err
	}
	if len(videoIdsDB) == 0 {
		return &rpc.VideoPublishResp{}, nil
	}

	videostr := make([]string, len(videoIdsDB))
	for i := range videoIdsDB {
		videostr[i] = strconv.Itoa(int(videoIdsDB[i].VideoId))
	}

	videosDB, err := l.gets.Gets(&video.VideoReq{VideoIds: strings.Join(videostr, ",")})
	if err != nil {
		return nil, err
	}
	return &rpc.VideoPublishResp{
		Videos: videosDB.Videos,
	}, nil
}
