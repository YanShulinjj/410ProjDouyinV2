package logic

import (
	"410proj/apps/comment/rpc/internal/svc"
	"410proj/apps/comment/rpc/model"
	"410proj/apps/comment/rpc/rpc"
	"context"

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

func (l *AddLogic) Add(in *rpc.AddCommentReq) (*rpc.AddCommentResp, error) {
	// 发布评论, 请在api层保证userid 和 videoid的准确性
	comment := model.Comment{
		UserId:  int64(in.UserId),
		VideoId: int64(in.VideoId),
		Content: in.Content,
	}
	res, err := l.svcCtx.CommentModel.Insert(l.ctx, &comment)
	if err != nil {
		return nil, err
	}
	commenID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	// 更新commen_num表
	// 2. 先获取点赞数目表
	commentNumDB, err := l.svcCtx.CommentNumModel.FindOneByVideoId(l.ctx, int64(in.VideoId))
	if err != nil {
		if err == model.ErrNotFound {
			// 新建条目
			commentnum := model.VideoCommentNum{
				VideoId:    int64(in.VideoId),
				CommentNum: 1,
			}
			_, err = l.svcCtx.CommentNumModel.Insert(l.ctx, &commentnum)
			if err != nil {
				return nil, err
			}
			return &rpc.AddCommentResp{}, nil
		}
		return nil, err
	}

	commentNumDB.CommentNum++
	// 再更新到DB
	err = l.svcCtx.CommentNumModel.Update(l.ctx, commentNumDB)
	if err != nil {
		return nil, err
	}
	return &rpc.AddCommentResp{
		CommentId: uint64(commenID),
	}, nil
}
