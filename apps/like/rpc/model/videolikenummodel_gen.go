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
	videoLikeNumFieldNames          = builder.RawFieldNames(&VideoLikeNum{})
	videoLikeNumRows                = strings.Join(videoLikeNumFieldNames, ",")
	videoLikeNumRowsExpectAutoSet   = strings.Join(stringx.Remove(videoLikeNumFieldNames, "`like_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	videoLikeNumRowsWithPlaceHolder = strings.Join(stringx.Remove(videoLikeNumFieldNames, "`like_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheVideoLikeNumLikeIdPrefix  = "cache:videoLikeNum:likeId:"
	cacheVideoLikeNumVideoIdPrefix = "cache:videoLikeNum:videoId:"
)

type (
	videoLikeNumModel interface {
		Insert(ctx context.Context, data *VideoLikeNum) (sql.Result, error)
		FindOne(ctx context.Context, likeId int64) (*VideoLikeNum, error)
		FindOneByVideoId(ctx context.Context, videoId int64) (*VideoLikeNum, error)
		Update(ctx context.Context, data *VideoLikeNum) error
		Delete(ctx context.Context, likeId int64) error
	}

	defaultVideoLikeNumModel struct {
		sqlc.CachedConn
		table string
	}

	VideoLikeNum struct {
		LikeId     int64     `db:"like_id"`     // 关系ID
		VideoId    int64     `db:"video_id"`    // 视频id
		Likes      int64     `db:"likes"`       // 视频的点赞数
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newVideoLikeNumModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideoLikeNumModel {
	return &defaultVideoLikeNumModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video_like_num`",
	}
}

func (m *defaultVideoLikeNumModel) Delete(ctx context.Context, likeId int64) error {
	data, err := m.FindOne(ctx, likeId)
	if err != nil {
		return err
	}

	videoLikeNumLikeIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumLikeIdPrefix, likeId)
	videoLikeNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumVideoIdPrefix, data.VideoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `like_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, likeId)
	}, videoLikeNumLikeIdKey, videoLikeNumVideoIdKey)
	return err
}

func (m *defaultVideoLikeNumModel) FindOne(ctx context.Context, likeId int64) (*VideoLikeNum, error) {
	videoLikeNumLikeIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumLikeIdPrefix, likeId)
	var resp VideoLikeNum
	err := m.QueryRowCtx(ctx, &resp, videoLikeNumLikeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `like_id` = ? limit 1", videoLikeNumRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, likeId)
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

func (m *defaultVideoLikeNumModel) FindOneByVideoId(ctx context.Context, videoId int64) (*VideoLikeNum, error) {
	videoLikeNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumVideoIdPrefix, videoId)
	var resp VideoLikeNum
	err := m.QueryRowIndexCtx(ctx, &resp, videoLikeNumVideoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videoLikeNumRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, videoId); err != nil {
			return nil, err
		}
		return resp.LikeId, nil
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

func (m *defaultVideoLikeNumModel) Insert(ctx context.Context, data *VideoLikeNum) (sql.Result, error) {
	videoLikeNumLikeIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumLikeIdPrefix, data.LikeId)
	videoLikeNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumVideoIdPrefix, data.VideoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, videoLikeNumRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.VideoId, data.Likes)
	}, videoLikeNumLikeIdKey, videoLikeNumVideoIdKey)
	return ret, err
}

func (m *defaultVideoLikeNumModel) Update(ctx context.Context, newData *VideoLikeNum) error {
	data, err := m.FindOne(ctx, newData.LikeId)
	if err != nil {
		return err
	}

	videoLikeNumLikeIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumLikeIdPrefix, data.LikeId)
	videoLikeNumVideoIdKey := fmt.Sprintf("%s%v", cacheVideoLikeNumVideoIdPrefix, data.VideoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `like_id` = ?", m.table, videoLikeNumRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.VideoId, newData.Likes, newData.LikeId)
	}, videoLikeNumLikeIdKey, videoLikeNumVideoIdKey)
	return err
}

func (m *defaultVideoLikeNumModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVideoLikeNumLikeIdPrefix, primary)
}

func (m *defaultVideoLikeNumModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `like_id` = ? limit 1", videoLikeNumRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoLikeNumModel) tableName() string {
	return m.table
}
