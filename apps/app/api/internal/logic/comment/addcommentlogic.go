package comment

import (
	"410proj/apps/app/api/internal/logic/user"
	"410proj/apps/comment/rpc/comment"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"context"
	"time"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	userApi *user.UserinfoLogic
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		userApi: user.NewUserinfoLogic(ctx, svcCtx),
	}
}

func (l *AddCommentLogic) AddComment(req *types.AddCommentReq) (resp *types.AddCommentResp, err error) {
	// action_type = 1 点赞
	// 通过token获取userid
	userId, err := jwtx.GetUserId(l.svcCtx.Config.JwtAuth.AccessSecret, req.Token)
	if err != nil {
		return &types.AddCommentResp{
			CommentResponse: types.CommentResponse{
				StatusCode: int32(xerr.ReuqestParamErr),
				Msg:        "评论失败",
			}}, nil
	}

	// 获取用户信息
	user, err := l.userApi.Userinfo(&types.UserInfoReq{UserId: userId})
	if err != nil {
		return
	}

	var commentId uint64
	if req.ActionType == 1 {
		commentRpcReq := comment.AddCommentReq{
			UserId:  userId,
			VideoId: req.VideoId,
			Content: req.Comment,
		}
		res, err := l.svcCtx.CommentRPC.Add(l.ctx, &commentRpcReq)
		if err != nil {
			return &types.AddCommentResp{
				CommentResponse: types.CommentResponse{
					StatusCode: int32(xerr.ReuqestParamErr),
					Msg:        "评论失败",
				}}, nil
		}
		commentId = res.CommentId
	} else if req.ActionType == 2 {
		// 取消点赞
		// TODO
		commentRpcReq := comment.DropCommentReq{
			CommentId: req.Comment_id,
			VideoId:   req.VideoId,
		}
		_, err := l.svcCtx.CommentRPC.Drop(l.ctx, &commentRpcReq)
		if err != nil {
			return &types.AddCommentResp{
				CommentResponse: types.CommentResponse{
					StatusCode: int32(xerr.ReuqestParamErr),
					Msg:        "删除失败",
				}}, nil
		}
	}
	commenInfo := types.CommentInfo{
		Id:         commentId,
		User:       user.UserInfo,
		Content:    req.Comment,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	return &types.AddCommentResp{
		CommentResponse: types.CommentResponse{
			StatusCode: 0,
			Msg:        "ok",
		},
		Comment: commenInfo,
	}, nil
}
