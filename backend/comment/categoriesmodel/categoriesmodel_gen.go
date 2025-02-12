// Code generated by goctl. DO NOT EDIT!

package categoriesmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"tiktok/common/globalkey"
)

var (
	categoriesFieldNames          = builder.RawFieldNames(&Categories{})
	categoriesRows                = strings.Join(categoriesFieldNames, ",")
	categoriesRowsExpectAutoSet   = strings.Join(stringx.Remove(categoriesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	categoriesRowsWithPlaceHolder = strings.Join(stringx.Remove(categoriesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTiktokCategoriesIdPrefix = "cache:tiktok:categories:id:"
)

type (
	categoriesModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Categories) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Categories, error)
		Update(ctx context.Context, session sqlx.Session, data *Categories) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Categories) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *Categories) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Categories, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Categories, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Categories, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Categories, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Categories, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultCategoriesModel struct {
		sqlc.CachedConn
		table string
	}

	Categories struct {
		Id         int64     `db:"id"`
		Vid        int64     `db:"vid"`
		Label      string    `db:"label"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
	}
)

func newCategoriesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCategoriesModel {
	return &defaultCategoriesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`categories`",
	}
}

func (m *defaultCategoriesModel) Insert(ctx context.Context, session sqlx.Session, data *Categories) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	tiktokCategoriesIdKey := fmt.Sprintf("%s%v", cacheTiktokCategoriesIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, categoriesRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Vid, data.Label, data.DeleteTime, data.DelState, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.Vid, data.Label, data.DeleteTime, data.DelState, data.Version)
	}, tiktokCategoriesIdKey)
}

func (m *defaultCategoriesModel) FindOne(ctx context.Context, id int64) (*Categories, error) {
	tiktokCategoriesIdKey := fmt.Sprintf("%s%v", cacheTiktokCategoriesIdPrefix, id)
	var resp Categories
	err := m.QueryRowCtx(ctx, &resp, tiktokCategoriesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", categoriesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
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

func (m *defaultCategoriesModel) Update(ctx context.Context, session sqlx.Session, data *Categories) (sql.Result, error) {
	tiktokCategoriesIdKey := fmt.Sprintf("%s%v", cacheTiktokCategoriesIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, categoriesRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Vid, data.Label, data.DeleteTime, data.DelState, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Vid, data.Label, data.DeleteTime, data.DelState, data.Version, data.Id)
	}, tiktokCategoriesIdKey)
}

func (m *defaultCategoriesModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *Categories) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "CategoriesModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultCategoriesModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultCategoriesModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultCategoriesModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Categories, error) {

	builder = builder.Columns(categoriesRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Categories
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCategoriesModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Categories, error) {

	builder = builder.Columns(categoriesRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Categories
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCategoriesModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Categories, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(categoriesRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*Categories
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultCategoriesModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Categories, error) {

	builder = builder.Columns(categoriesRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Categories
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCategoriesModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Categories, error) {

	builder = builder.Columns(categoriesRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Categories
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCategoriesModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultCategoriesModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultCategoriesModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Categories) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	tiktokCategoriesIdKey := fmt.Sprintf("%s%v", cacheTiktokCategoriesIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, categoriesRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Vid, data.Label, data.DeleteTime, data.DelState, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.Vid, data.Label, data.DeleteTime, data.DelState, data.Version, data.Id, oldVersion)
	}, tiktokCategoriesIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}
func (m *defaultCategoriesModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	tiktokCategoriesIdKey := fmt.Sprintf("%s%v", cacheTiktokCategoriesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, tiktokCategoriesIdKey)
	return err
}
func (m *defaultCategoriesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokCategoriesIdPrefix, primary)
}
func (m *defaultCategoriesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", categoriesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultCategoriesModel) tableName() string {
	return m.table
}
