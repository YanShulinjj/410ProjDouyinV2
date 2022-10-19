package svc

import (
	"410proj/apps/videos/rpc/internal/config"
	"410proj/apps/videos/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 添加 video model 依赖
	VideoModel   model.VideoModel
	PublishModel model.UserPublishMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		VideoModel: model.NewVideoModel(sqlConn,
			c.CacheRedis),
		PublishModel: model.NewUserPublishMapModel(sqlConn,
			c.CacheRedis),
	}
}
