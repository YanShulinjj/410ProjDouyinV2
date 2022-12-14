// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	videoCommentNumFieldNames          = builder.RawFieldNames(&VideoCommentNum{})
	videoCommentNumRows                = strings.Join(videoCommentNumFieldNames, ",")
	videoCommentNumRowsExpectAutoSet   = strings.Join(stringx.Remove(videoCommentNumFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	videoCommentNumRowsWithPlaceHolder = strings.Join(stringx.Remove(videoCommentNumFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheVideoCommentNumIdPrefix      = "cache:videoCommentNum:id:"
	cacheVideoCommentNumVideoIdPrefix = "cache:videoCommentNum:videoId:"
)

type (
	videoCommentNumModel interface {
		Insert(ctx context.Context, data *VideoCommentNum) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*VideoCommentNum, error)
		FindOneByVideoId(ctx context.Context, videoId int64) (*VideoCommentNum, error)
		Update(ctx context.Context, data *VideoCommentNum) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoCommentNumModel struct {
		sqlc.CachedConn
		table string
	}

	VideoCommentNum struct {
		Id         int64     `db:"id"`          // ID
		VideoId    int64     `db:"video_id"`    // 视频id
		CommentNum int64     `db:"comment_num"` // 视频的点赞数
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newVideoCommentNumModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideoCommentNumModel {
	return &defaultVideoCommentNumModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video_comment_num`",
	}
}

func (m *defaultVideoCommentNumModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	videoCommentNumIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumIdPrefix, id)
	videoCommentNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumVideoIdPrefix, data.VideoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, videoCommentNumIdKey, videoCommentNumVideoIdKey)
	return err
}

func (m *defaultVideoCommentNumModel) FindOne(ctx context.Context, id int64) (*VideoCommentNum, error) {
	videoCommentNumIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumIdPrefix, id)
	var resp VideoCommentNum
	err := m.QueryRowCtx(ctx, &resp, videoCommentNumIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoCommentNumRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoCommentNumModel) FindOneByVideoId(ctx context.Context, videoId int64) (*VideoCommentNum, error) {
	videoCommentNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumVideoIdPrefix, videoId)
	var resp VideoCommentNum
	err := m.QueryRowIndexCtx(ctx, &resp, videoCommentNumVideoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videoCommentNumRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, videoId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoCommentNumModel) Insert(ctx context.Context, data *VideoCommentNum) (sql.Result, error) {
	videoCommentNumIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumIdPrefix, data.Id)
	videoCommentNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumVideoIdPrefix, data.VideoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, videoCommentNumRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.VideoId, data.CommentNum)
	}, videoCommentNumIdKey, videoCommentNumVideoIdKey)
	return ret, err
}

func (m *defaultVideoCommentNumModel) Update(ctx context.Context, newData *VideoCommentNum) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	videoCommentNumIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumIdPrefix, data.Id)
	videoCommentNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoCommentNumVideoIdPrefix, data.VideoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoCommentNumRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.VideoId, newData.CommentNum, newData.Id)
	}, videoCommentNumIdKey, videoCommentNumVideoIdKey)
	return err
}

func (m *defaultVideoCommentNumModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVideoCommentNumIdPrefix, primary)
}

func (m *defaultVideoCommentNumModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoCommentNumRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoCommentNumModel) tableName() string {
	return m.table
}
