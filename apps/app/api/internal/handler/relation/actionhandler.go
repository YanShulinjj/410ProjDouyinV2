package relation

import (
	"net/http"

	"410proj/apps/app/api/internal/logic/relation"
	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionRelationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := relation.NewActionLogic(r.Context(), svcCtx)
		resp, err := l.Action(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
