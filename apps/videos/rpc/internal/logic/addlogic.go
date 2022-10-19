package logic

import (
	"410proj/apps/videos/rpc/model"
	"410proj/pkg/xerr"
	"context"
	"github.com/pkg/errors"

	"410proj/apps/videos/rpc/internal/svc"
	"410proj/apps/videos/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *rpc.AddVideoReq) (*rpc.AddVideoResp, error) {
	// 发布视频
	if len(in.PlayUrl) == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ReuqestParamErr),
			"发布视频失败!")
	}
	video := model.Video{
		UserId:   in.AuthorId,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
	}
	res, err := l.svcCtx.VideoModel.Insert(l.ctx, &video)
	if err != nil {
		return nil, err
	}
	// 添加到user_publish_map
	vid, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ReuqestParamErr),
			"获取vid失败!")
	}
	item := model.UserPublishMap{
		UserId:  int64(in.AuthorId),
		VideoId: vid,
	}
	_, err = l.svcCtx.PublishModel.Insert(l.ctx, &item)
	if err != nil {
		return nil, err
	}

	return &rpc.AddVideoResp{}, nil
}
