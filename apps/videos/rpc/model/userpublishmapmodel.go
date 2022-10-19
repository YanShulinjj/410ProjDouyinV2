package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserPublishMapModel = (*customUserPublishMapModel)(nil)

type (
	// UserPublishMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPublishMapModel.
	UserPublishMapModel interface {
		userPublishMapModel
	}

	customUserPublishMapModel struct {
		*defaultUserPublishMapModel
	}
)

// NewUserPublishMapModel returns a model for the database table.
func NewUserPublishMapModel(conn sqlx.SqlConn, c cache.CacheConf) UserPublishMapModel {
	return &customUserPublishMapModel{
		defaultUserPublishMapModel: newUserPublishMapModel(conn, c),
	}
}
