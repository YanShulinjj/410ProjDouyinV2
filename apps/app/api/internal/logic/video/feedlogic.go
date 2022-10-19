package video

import (
	"410proj/apps/app/api/internal/logic/user"
	"410proj/apps/comment/rpc/comment"
	"410proj/apps/like/rpc/like"
	"410proj/apps/videos/rpc/rpc"
	"410proj/apps/videos/rpc/video"
	"410proj/pkg/fmtx"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"context"
	"fmt"
	"sync"
	"time"

	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	userApi *user.UserinfoLogic
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		userApi: user.NewUserinfoLogic(ctx, svcCtx),
	}
}

func (l *FeedLogic) Feed(req *types.VideoFeedReq) (resp *types.VideoFeedResp, err error) {
	var in video.VideoFeedReq
	if req.Latest_time == 0 {
		in.LatestTime = time.Now().Unix()
	} else {
		in.LatestTime = req.Latest_time
	}
	// 如果携带token，获取token中的user_id
	var user_id uint64
	if len(req.Token) > 0 {
		user_id, _ = jwtx.GetUserId(l.svcCtx.Config.JwtAuth.AccessSecret, req.Token)
	}
	res, err := l.svcCtx.VideoRPC.Feeds(l.ctx, &in)
	if err != nil {
		return &types.VideoFeedResp{
			VideoResponse: types.VideoResponse{
				StatusCode: int32(xerr.ReuqestParamErr),
				Msg:        "获取视频流失败：" + err.Error(),
			},
		}, nil
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
				fmt.Println("1.*********:", err)
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
			likeRpcReq := like.IsLikeReq{UserId: user_id, VideoId: v.Id}
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

	return &types.VideoFeedResp{
		VideoResponse: types.VideoResponse{
			StatusCode: 0,
			Msg:        "ok",
		},
		Videos: respVideos,
	}, nil
}
