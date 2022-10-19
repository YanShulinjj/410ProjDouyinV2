package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoLikeNumModel = (*customVideoLikeNumModel)(nil)

type (
	// VideoLikeNumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoLikeNumModel.
	VideoLikeNumModel interface {
		videoLikeNumModel
	}

	customVideoLikeNumModel struct {
		*defaultVideoLikeNumModel
	}
)

// NewVideoLikeNumModel returns a model for the database table.
func NewVideoLikeNumModel(conn sqlx.SqlConn, c cache.CacheConf) VideoLikeNumModel {
	return &customVideoLikeNumModel{
		defaultVideoLikeNumModel: newVideoLikeNumModel(conn, c),
	}
}
