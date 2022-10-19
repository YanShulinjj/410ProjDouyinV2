package favorite

import (
	"410proj/apps/app/api/internal/logic/user"
	"410proj/apps/comment/rpc/comment"
	"410proj/apps/like/rpc/like"
	"410proj/apps/videos/rpc/rpc"
	"410proj/apps/videos/rpc/video"
	"410proj/pkg/fmtx"
	"context"
	"sync"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikelistLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	userApi *user.UserinfoLogic
}

func NewLikelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikelistLogic {
	return &LikelistLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		userApi: user.NewUserinfoLogic(ctx, svcCtx),
	}
}

func (l *LikelistLogic) Likelist(req *types.VideoListReq) (resp *types.LikeVideoListResp, err error) {
	// todo: add your logic here and delete this line
	likeVideoIdsReq := like.User{
		UserId: req.UserId,
	}
	likeVideoIdsResp, err := l.svcCtx.LikeRPC.GetLikeList(
		l.ctx, &likeVideoIdsReq)
	if err != nil {
		return nil, err
	}
	if len(likeVideoIdsResp.VideoIds) == 0 {
		return &types.LikeVideoListResp{
			VideoResponse: types.VideoResponse{
				StatusCode: 0,
				Msg:        "ok",
			},
		}, nil
	}
	likeVideosReq := video.VideoReq{
		VideoIds: likeVideoIdsResp.VideoIds,
	}
	res, err := l.svcCtx.VideoRPC.Gets(
		l.ctx, &likeVideosReq)
	if err != nil {
		return nil, err
	}

	respVideos := make([]*types.VideoItem, len(res.Videos))
	// 此处采用多线程获取信息
	wg := sync.WaitGroup{}
	for i, video := range res.Videos {
		wg.Add(1)
		go func(j int, v *rpc.VideoItem) {
			defer wg.Done()
			// 通过userid 获取用户信息
			user, err := l.userApi.Userinfo(&types.UserInfoReq{UserId: video.AuthorId})
			if err != nil {
				return
			}
			// 通过 LikeRpc 获取该视频的点赞数
			likeNumRpcReq := like.LikeNumReq{VideoId: v.Id}
			likeNumRpcResp, err := l.svcCtx.LikeRPC.GetLikeNum(l.ctx, &likeNumRpcReq)
			likenum := "0"
			if err == nil {
				likenum = fmtx.ItoA(likeNumRpcResp.Nums)
			}
			// 通过 CommentRpc 获取该视频的评论数
			commentRpcReq := comment.CommentNumReq{VideoId: v.Id}
			commentRpcResp, err := l.svcCtx.CommentRPC.GetCommentNum(l.ctx, &commentRpcReq)
			commentnum := "0"
			if err == nil {
				commentnum = fmtx.ItoA(commentRpcResp.Nums)
			}
			// 通过LikeRpc 获取该用户是否对此视频已点赞
			likeRpcReq := like.IsLikeReq{UserId: req.UserId, VideoId: v.Id}
			likeRpcResp, err := l.svcCtx.LikeRPC.IsLike(l.ctx, &likeRpcReq)
			if err != nil {
				return
			}
			respVideos[j] = &types.VideoItem{
				Id:            v.Id,
				User:          user.UserInfo,
				PlayURL:       v.PlayUrl,
				CoverURL:      v.CoverUrl,
				FavoriteCount: likenum,
				CommentCount:  commentnum,
				IsLike:        likeRpcResp.IsLike,
			}
		}(i, video)
	}
	wg.Wait()

	return &types.LikeVideoListResp{
		VideoResponse: types.VideoResponse{
			StatusCode: 0,
			Msg:        "ok",
		},
		Videos: respVideos,
	}, nil

	return
}
