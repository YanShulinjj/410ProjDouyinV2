package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoCommentNumModel = (*customVideoCommentNumModel)(nil)

type (
	// VideoCommentNumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoCommentNumModel.
	VideoCommentNumModel interface {
		videoCommentNumModel
	}

	customVideoCommentNumModel struct {
		*defaultVideoCommentNumModel
	}
)

// NewVideoCommentNumModel returns a model for the database table.
func NewVideoCommentNumModel(conn sqlx.SqlConn, c cache.CacheConf) VideoCommentNumModel {
	return &customVideoCommentNumModel{
		defaultVideoCommentNumModel: newVideoCommentNumModel(conn, c),
	}
}
