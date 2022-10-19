package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"410proj/apps/videos/rpc/internal/svc"
	"410proj/apps/videos/rpc/rpc"

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

func (l *GetsLogic) Gets(in *rpc.VideoReq) (*rpc.VideoResp, error) {
	// 请求的video_ids 类似于 1,2,3,4
	// 通过多个uid 获取用户列表
	vidstrs := strings.Split(in.VideoIds, ",")
	videos := []*rpc.VideoItem{}
	for _, vidstr := range vidstrs {
		fmt.Println("VVVVVVVVV", in.VideoIds, vidstrs)
		vid, err := strconv.Atoi(vidstr)
		if err != nil {
			return nil, err
		}
		videoDB, err := l.svcCtx.VideoModel.FindOne(l.ctx, uint64(vid))
		if err != nil {
			return nil, err
		}
		videoitem := &rpc.VideoItem{
			Id:         videoDB.VideoId,
			AuthorId:   videoDB.UserId,
			PlayUrl:    videoDB.PlayUrl,
			CoverUrl:   videoDB.CoverUrl,
			CreateTime: videoDB.CreateTime.UnixNano(),
			// Password: userDB.Password,
		}
		videos = append(videos, videoitem)
	}
	return &rpc.VideoResp{
		Videos: videos,
	}, nil
	return &rpc.VideoResp{}, nil
}
