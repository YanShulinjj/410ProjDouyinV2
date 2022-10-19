package video

import (
	"410proj/apps/app/api/internal/types"
	"net/http"

	"410proj/apps/app/api/internal/logic/video"
	"410proj/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoFeedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := video.NewFeedLogic(r.Context(), svcCtx)
		resp, err := l.Feed(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
