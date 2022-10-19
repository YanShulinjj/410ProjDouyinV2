package logic

import (
	"context"
	"strconv"
	"strings"

	"410proj/apps/like/rpc/internal/svc"
	"410proj/apps/like/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikeListLogic {
	return &GetLikeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikeListLogic) GetLikeList(in *rpc.User) (*rpc.LikeList, error) {
	// todo: add your logic here and delete this line
	videosDB, err := l.svcCtx.LikeModel.FindMany(l.ctx, int64(in.UserId))
	if err != nil {
		return nil, err
	}
	videostr := make([]string, len(videosDB))
	for i := range videosDB {
		videostr[i] = strconv.Itoa(int(videosDB[i].VideoId))
	}

	return &rpc.LikeList{
		VideoIds: strings.Join(videostr, ","),
	}, nil
}
