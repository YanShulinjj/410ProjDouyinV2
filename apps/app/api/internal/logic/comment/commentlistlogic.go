package comment

import (
	"410proj/apps/app/api/internal/logic/user"
	"410proj/apps/comment/rpc/comment"
	rpc2 "410proj/apps/comment/rpc/rpc"
	"410proj/pkg/xerr"
	"context"
	"fmt"
	"sync"
	"time"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	userApi *user.UserinfoLogic
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		userApi: user.NewUserinfoLogic(ctx, svcCtx),
	}
}

func (l *CommentListLogic) CommentList(req *types.VideoCommentListReq) (resp *types.VideoCommentListResp, err error) {
	// todo: add your logic here and delete this line
	commentReq := comment.CommentReq{
		VideoId: req.VideoId,
	}
	res, err := l.svcCtx.CommentRPC.Gets(l.ctx, &commentReq)
	if err != nil {
		return &types.VideoCommentListResp{
			CommentResponse: types.CommentResponse{
				StatusCode: int32(xerr.ReuqestParamErr),
				Msg:        "获取评论列表失败 ！",
			},
		}, err
	}

	comments := make([]*types.CommentInfo, len(res.Comments))

	// 此处采用多线程获取信息
	wg := sync.WaitGroup{}
	for i, comment := range res.Comments {
		wg.Add(1)
		go func(j int, v *rpc2.CommentItem) {
			defer wg.Done()
			// 通过userid 获取用户信息
			user, err := l.userApi.Userinfo(&types.UserInfoReq{UserId: v.UserId})
			if err != nil {
				fmt.Println("1.*********:", err)
				return
			}
			tm := time.Unix(v.CreateTime, 0)
			commenItem := &types.CommentInfo{
				Id:         v.CommentId,
				User:       user.UserInfo,
				Content:    v.Content,
				CreateDate: tm.Format("2006-01-02 15:04:05"),
			}
			comments[j] = commenItem
		}(i, comment)
	}
	wg.Wait()
	return &types.VideoCommentListResp{
		CommentResponse: types.CommentResponse{
			StatusCode: 0,
			Msg:        "ok",
		},
		Comments: comments,
	}, err
}
