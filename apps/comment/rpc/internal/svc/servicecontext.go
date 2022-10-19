package svc

import (
	"410proj/apps/comment/rpc/internal/config"
	"410proj/apps/comment/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 添加comment model 依赖
	CommentModel    model.CommentModel
	CommentNumModel model.VideoCommentNumModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		CommentModel:    model.NewCommentModel(sqlConn, c.CacheRedis),
		CommentNumModel: model.NewVideoCommentNumModel(sqlConn, c.CacheRedis),
	}
}
