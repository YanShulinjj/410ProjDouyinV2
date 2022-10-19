package svc

import (
	"410proj/apps/app/api/internal/config"
	"410proj/apps/app/api/internal/middleware"
	"410proj/apps/comment/rpc/comment"
	"410proj/apps/like/rpc/like"
	"410proj/apps/relation/rpc/relation"
	"410proj/apps/user/rpc/user"
	"410proj/apps/videos/rpc/video"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
)

type ServiceContext struct {
	Config      config.Config
	UserRPC     user.User
	VideoRPC    video.Video
	CommentRPC  comment.Comment
	LikeRPC     like.Like
	RelationRPC relation.Relation
	Jwt         rest.Middleware
	OssClient   *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	oc, err := oss.New(c.OSSEndpoint, c.AccessKeyID, c.AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:      c,
		UserRPC:     user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		VideoRPC:    video.NewVideo(zrpc.MustNewClient(c.VideoRPC)),
		CommentRPC:  comment.NewComment(zrpc.MustNewClient(c.CommentRPC)),
		LikeRPC:     like.NewLike(zrpc.MustNewClient(c.LikeRPC)),
		RelationRPC: relation.NewRelation(zrpc.MustNewClient(c.RelationRPC)),
		Jwt:         middleware.NewJwtMiddleware(c.JwtAuth.AccessSecret).Handle,
		OssClient:   oc,
	}
}
