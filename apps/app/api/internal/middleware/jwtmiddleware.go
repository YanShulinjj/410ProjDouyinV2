package middleware

import (
	"410proj/apps/app/api/internal/types"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type JwtMiddleware struct {
	secretKey string
}

func NewJwtMiddleware(secretKey string) *JwtMiddleware {
	return &JwtMiddleware{
		secretKey: secretKey,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("jwt middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		token := r.FormValue("token")
		if len(token) == 0 {
			token = r.PostForm.Get("token")
		}
		_, err := jwtx.IsValid(m.secretKey, token)
		if err != nil {
			httpx.OkJson(w, &types.UserResponse{
				StatusCode: int32(xerr.TokenNotMatchErr),
				Msg:        "token失效, 需要重新登陆",
			})
		} else {
			next(w, r)
		}
	}
}
