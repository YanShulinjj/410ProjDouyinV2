package comment

import (
	"net/http"

	"410proj/apps/app/api/internal/logic/comment"
	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewAddCommentLogic(r.Context(), svcCtx)
		resp, err := l.AddComment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
