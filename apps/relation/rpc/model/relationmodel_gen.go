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
	relationFieldNames          = builder.RawFieldNames(&Relation{})
	relationRows                = strings.Join(relationFieldNames, ",")
	relationRowsExpectAutoSet   = strings.Join(stringx.Remove(relationFieldNames, "`relation_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	relationRowsWithPlaceHolder = strings.Join(stringx.Remove(relationFieldNames, "`relation_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheRelationRelationIdPrefix = "cache:relation:relationId:"
	cacheRelationUserIdTypePrefix = "cache:relation:userId:type:"
)

type (
	relationModel interface {
		Insert(ctx context.Context, data *Relation) (sql.Result, error)
		FindOne(ctx context.Context, relationId uint64) (*Relation, error)
		FindOneByUserIdType(ctx context.Context, userId uint64, tp bool) (*Relation, error)
		Update(ctx context.Context, data *Relation) error
		Delete(ctx context.Context, relationId uint64) error
	}

	defaultRelationModel struct {
		sqlc.CachedConn
		table string
	}

	Relation struct {
		RelationId uint64    `db:"relation_id"` // 关系ID
		UserId     uint64    `db:"user_id"`     // 作者id
		ToUserIds  string    `db:"to_user_ids"` // 关系用户的id，形如 1,2,3
		Type       bool      `db:"type"`        // 0-follow, 1-follower
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newRelationModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRelationModel {
	return &defaultRelationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`relation`",
	}
}

func (m *defaultRelationModel) Delete(ctx context.Context, relationId uint64) error {
	data, err := m.FindOne(ctx, relationId)
	if err != nil {
		return err
	}

	relationRelationIdKey := fmt.Sprintf("%s%v", cacheRelationRelationIdPrefix, relationId)
	relationUserIdTypeKey := fmt.Sprintf("%s%v:%v", cacheRelationUserIdTypePrefix, data.UserId, data.Type)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `relation_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, relationId)
	}, relationRelationIdKey, relationUserIdTypeKey)
	return err
}

func (m *defaultRelationModel) FindOne(ctx context.Context, relationId uint64) (*Relation, error) {
	relationRelationIdKey := fmt.Sprintf("%s%v", cacheRelationRelationIdPrefix, relationId)
	var resp Relation
	err := m.QueryRowCtx(ctx, &resp, relationRelationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `relation_id` = ? limit 1", relationRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, relationId)
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

func (m *defaultRelationModel) FindOneByUserIdType(ctx context.Context, userId uint64, tp bool) (*Relation, error) {
	relationUserIdTypeKey := fmt.Sprintf("%s%v:%v", cacheRelationUserIdTypePrefix, userId, tp)
	var resp Relation
	err := m.QueryRowIndexCtx(ctx, &resp, relationUserIdTypeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `type` = ? limit 1", relationRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, tp); err != nil {
			return nil, err
		}
		return resp.RelationId, nil
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

func (m *defaultRelationModel) Insert(ctx context.Context, data *Relation) (sql.Result, error) {
	relationRelationIdKey := fmt.Sprintf("%s%v", cacheRelationRelationIdPrefix, data.RelationId)
	relationUserIdTypeKey := fmt.Sprintf("%s%v:%v", cacheRelationUserIdTypePrefix, data.UserId, data.Type)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, relationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserIds, data.Type)
	}, relationRelationIdKey, relationUserIdTypeKey)
	return ret, err
}

func (m *defaultRelationModel) Update(ctx context.Context, newData *Relation) error {
	data, err := m.FindOne(ctx, newData.RelationId)
	if err != nil {
		return err
	}

	relationRelationIdKey := fmt.Sprintf("%s%v", cacheRelationRelationIdPrefix, data.RelationId)
	relationUserIdTypeKey := fmt.Sprintf("%s%v:%v", cacheRelationUserIdTypePrefix, data.UserId, data.Type)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `relation_id` = ?", m.table, relationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.ToUserIds, newData.Type, newData.RelationId)
	}, relationRelationIdKey, relationUserIdTypeKey)
	return err
}

func (m *defaultRelationModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRelationRelationIdPrefix, primary)
}

func (m *defaultRelationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `relation_id` = ? limit 1", relationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRelationModel) tableName() string {
	return m.table
}
