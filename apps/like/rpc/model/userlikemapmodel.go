package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLikeMapModel = (*customUserLikeMapModel)(nil)

type (
	// UserLikeMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLikeMapModel.
	UserLikeMapModel interface {
		userLikeMapModel
	}

	customUserLikeMapModel struct {
		*defaultUserLikeMapModel
	}
)

// NewUserLikeMapModel returns a model for the database table.
func NewUserLikeMapModel(conn sqlx.SqlConn, c cache.CacheConf) UserLikeMapModel {
	return &customUserLikeMapModel{
		defaultUserLikeMapModel: newUserLikeMapModel(conn, c),
	}
}
