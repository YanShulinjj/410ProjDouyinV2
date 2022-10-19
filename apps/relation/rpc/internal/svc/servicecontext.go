package svc

import (
	"410proj/apps/relation/rpc/internal/config"
	"410proj/apps/relation/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 添加 relation model 依赖
	RelationModel model.RelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RelationModel: model.NewRelationModel(sqlConn,
			c.CacheRedis),
	}
}
