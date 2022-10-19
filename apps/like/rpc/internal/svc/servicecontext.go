package svc

import (
	"410proj/apps/like/rpc/internal/config"
	"410proj/apps/like/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 添加 like model 依赖
	LikeNumModel model.VideoLikeNumModel
	LikeModel    model.UserLikeMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		LikeNumModel: model.NewVideoLikeNumModel(sqlConn, c.CacheRedis),
		LikeModel:    model.NewUserLikeMapModel(sqlConn, c.CacheRedis),
	}
}
