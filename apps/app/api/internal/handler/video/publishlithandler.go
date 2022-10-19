package video

import (
	"net/http"

	"410proj/apps/app/api/internal/logic/video"
	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishlitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := video.NewPublishlitLogic(r.Context(), svcCtx)
		resp, err := l.Publishlit(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
