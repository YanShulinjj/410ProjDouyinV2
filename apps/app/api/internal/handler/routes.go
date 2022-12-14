// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"410proj/apps/app/api/internal/handler/file"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"

	comment "410proj/apps/app/api/internal/handler/comment"
	favorite "410proj/apps/app/api/internal/handler/favorite"
	relation "410proj/apps/app/api/internal/handler/relation"
	user "410proj/apps/app/api/internal/handler/user"
	video "410proj/apps/app/api/internal/handler/video"
	"410proj/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: user.UserinfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: comment.CommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/comment"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: comment.AddCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/comment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/feed",
				Handler: video.FeedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: video.AddHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: video.PublishlitHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: favorite.LikeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: favorite.LikelistHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/favorite"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: relation.ActionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follow/list",
					Handler: relation.FollowlistHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follower/list",
					Handler: relation.FollowerlistHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/relation"),
	)

	// ????????????
	dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
	patern := "/data/"
	dirpath := "./data/"
	prefix := "/"
	for i := 1; i < len(dirlevel); i++ {
		path := prefix + strings.Join(dirlevel[:i], "/")
		//???????????? /asset
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: file.Dirhandler(patern, dirpath),
			})
		logx.Infof("register dir  %s  %s", path,dirpath)
	}
}
