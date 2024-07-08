// Code generated by goctl. DO NOT EDIT.

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
	noteFieldNames          = builder.RawFieldNames(&Note{})
	noteRows                = strings.Join(noteFieldNames, ",")
	noteRowsExpectAutoSet   = strings.Join(stringx.Remove(noteFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	noteRowsWithPlaceHolder = strings.Join(stringx.Remove(noteFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheNoteIdPrefix = "cache:note:id:"
)

type (
	noteModel interface {
		Insert(ctx context.Context, data *Note) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Note, error)
		Update(ctx context.Context, data *Note) error
		Delete(ctx context.Context, id int64) error
	}

	defaultNoteModel struct {
		sqlc.CachedConn
		table string
	}

	Note struct {
		Id        int64          `db:"id"`
		Title     sql.NullString `db:"title"`
		FileId    sql.NullInt64  `db:"file_id"`
		CreatedAt time.Time      `db:"created_at"`
		CreatorId int64          `db:"creator_id"`
		IsDeleted int64          `db:"is_deleted"`
	}
)

func newNoteModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultNoteModel {
	return &defaultNoteModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`note`",
	}
}

func (m *defaultNoteModel) Delete(ctx context.Context, id int64) error {
	noteIdKey := fmt.Sprintf("%s%v", cacheNoteIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, noteIdKey)
	return err
}

func (m *defaultNoteModel) FindOne(ctx context.Context, id int64) (*Note, error) {
	noteIdKey := fmt.Sprintf("%s%v", cacheNoteIdPrefix, id)
	var resp Note
	err := m.QueryRowCtx(ctx, &resp, noteIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", noteRows, m.table)
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

func (m *defaultNoteModel) Insert(ctx context.Context, data *Note) (sql.Result, error) {
	noteIdKey := fmt.Sprintf("%s%v", cacheNoteIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, noteRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.FileId, data.CreatorId, data.IsDeleted)
	}, noteIdKey)
	return ret, err
}

func (m *defaultNoteModel) Update(ctx context.Context, data *Note) error {
	noteIdKey := fmt.Sprintf("%s%v", cacheNoteIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, noteRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.FileId, data.CreatorId, data.IsDeleted, data.Id)
	}, noteIdKey)
	return err
}

func (m *defaultNoteModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheNoteIdPrefix, primary)
}

func (m *defaultNoteModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", noteRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultNoteModel) tableName() string {
	return m.table
}