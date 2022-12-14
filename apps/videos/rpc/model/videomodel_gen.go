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
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`video_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`video_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheVideoVideoIdPrefix  = "cache:video:videoId:"
	cacheVideoVideoIdsPrefix = "cache:video:videoIds:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, videoId uint64) (*Video, error)
		FindManyBefore(ctx context.Context, time string, N int) ([]*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, videoId uint64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		VideoId    uint64    `db:"video_id"`    // 视频ID
		UserId     uint64    `db:"user_id"`     // 视频作者id
		PlayUrl    string    `db:"play_url"`    // 播放外链
		CoverUrl   string    `db:"cover_url"`   // 封面外链
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, videoId uint64) error {
	videoVideoIdKey := fmt.Sprintf("%s%v", cacheVideoVideoIdPrefix, videoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `video_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, videoId)
	}, videoVideoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, videoId uint64) (*Video, error) {
	videoVideoIdKey := fmt.Sprintf("%s%v", cacheVideoVideoIdPrefix, videoId)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, videoVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, videoId)
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

func (m *defaultVideoModel) FindManyBefore(ctx context.Context, time string, N int) ([]*Video, error) {
	var resp []*Video
	query := fmt.Sprintf("select %s from %s where `create_time` < '%s' order by `create_time` desc limit %d", videoRows, m.table, time, N)

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	videoVideoIdKey := fmt.Sprintf("%s%v", cacheVideoVideoIdPrefix, data.VideoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.PlayUrl, data.CoverUrl)
	}, videoVideoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	videoVideoIdKey := fmt.Sprintf("%s%v", cacheVideoVideoIdPrefix, data.VideoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `video_id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PlayUrl, data.CoverUrl, data.VideoId)
	}, videoVideoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVideoVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
