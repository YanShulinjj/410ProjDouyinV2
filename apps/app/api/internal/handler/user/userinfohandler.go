package user

import (
	"net/http"

	"410proj/apps/app/api/internal/logic/user"
	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
