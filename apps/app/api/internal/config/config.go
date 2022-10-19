package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRPC     zrpc.RpcClientConf
	VideoRPC    zrpc.RpcClientConf
	LikeRPC     zrpc.RpcClientConf
	RelationRPC zrpc.RpcClientConf
	CommentRPC  zrpc.RpcClientConf

	// 阿里云oss
	UseOSS          bool
	OSSEndpoint     string
	AccessKeyID     string
	AccessKeySecret string

	//
	FfmpegExecPath string

	ServerPath string
}
