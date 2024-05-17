// Code generated by goctl. DO NOT EDIT!

package model

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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTiktokUserIdPrefix       = "cache:tiktok:user:id:"
	cacheTiktokUserEmailPrefix    = "cache:tiktok:user:email:"
	cacheTiktokUserMobilePrefix   = "cache:tiktok:user:mobile:"
	cacheTiktokUserUsernamePrefix = "cache:tiktok:user:username:"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *User) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*User, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*User, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id            int64          `db:"id"`
		Username      string         `db:"username"` // 用户账号
		Email         string         `db:"email"`    // 用户邮件
		Nickname      string         `db:"nickname"` // 用户昵称
		Gender        int64          `db:"gender"`   // 用户性别
		Role          string         `db:"role"`     // 角色
		Status        int64          `db:"status"`   // 用户状态
		Mobile        string         `db:"mobile"`   // 用户电话
		Password      string         `db:"password"` // 用户密码
		Dec           string         `db:"dec"`      // 个性签名
		Avatar        string         `db:"avatar"`   // 头像
		CreateTime    time.Time      `db:"create_time"`
		UpdateTime    time.Time      `db:"update_time"`
		DeleteTime    time.Time      `db:"delete_time"`
		DelState      int64          `db:"del_state"`
		Version       int64          `db:"version"`        // 版本号
		BackgroundUrl sql.NullString `db:"background_url"` // 用户主页背景
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	tiktokUserEmailKey := fmt.Sprintf("%s%v", cacheTiktokUserEmailPrefix, data.Email)
	tiktokUserIdKey := fmt.Sprintf("%s%v", cacheTiktokUserIdPrefix, data.Id)
	tiktokUserMobileKey := fmt.Sprintf("%s%v", cacheTiktokUserMobilePrefix, data.Mobile)
	tiktokUserUsernameKey := fmt.Sprintf("%s%v", cacheTiktokUserUsernamePrefix, data.Username)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Username, data.Email, data.Nickname, data.Gender, data.Role, data.Status, data.Mobile, data.Password, data.Dec, data.Avatar, data.DeleteTime, data.DelState, data.Version, data.BackgroundUrl)
		}
		return conn.ExecCtx(ctx, query, data.Username, data.Email, data.Nickname, data.Gender, data.Role, data.Status, data.Mobile, data.Password, data.Dec, data.Avatar, data.DeleteTime, data.DelState, data.Version, data.BackgroundUrl)
	}, tiktokUserEmailKey, tiktokUserIdKey, tiktokUserMobileKey, tiktokUserUsernameKey)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	tiktokUserIdKey := fmt.Sprintf("%s%v", cacheTiktokUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, tiktokUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	tiktokUserEmailKey := fmt.Sprintf("%s%v", cacheTiktokUserEmailPrefix, email)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, tiktokUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? and del_state = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email, globalkey.DelStateNo); err != nil {
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

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	tiktokUserMobileKey := fmt.Sprintf("%s%v", cacheTiktokUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, tiktokUserMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? and del_state = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile, globalkey.DelStateNo); err != nil {
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

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	tiktokUserUsernameKey := fmt.Sprintf("%s%v", cacheTiktokUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, tiktokUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? and del_state = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username, globalkey.DelStateNo); err != nil {
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

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, newData *User) (sql.Result, error) {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return nil, err
	}
	tiktokUserEmailKey := fmt.Sprintf("%s%v", cacheTiktokUserEmailPrefix, data.Email)
	tiktokUserIdKey := fmt.Sprintf("%s%v", cacheTiktokUserIdPrefix, data.Id)
	tiktokUserMobileKey := fmt.Sprintf("%s%v", cacheTiktokUserMobilePrefix, data.Mobile)
	tiktokUserUsernameKey := fmt.Sprintf("%s%v", cacheTiktokUserUsernamePrefix, data.Username)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.Username, newData.Email, newData.Nickname, newData.Gender, newData.Role, newData.Status, newData.Mobile, newData.Password, newData.Dec, newData.Avatar, newData.DeleteTime, newData.DelState, newData.Version, newData.BackgroundUrl, newData.Id)
		}
		return conn.ExecCtx(ctx, query, newData.Username, newData.Email, newData.Nickname, newData.Gender, newData.Role, newData.Status, newData.Mobile, newData.Password, newData.Dec, newData.Avatar, newData.DeleteTime, newData.DelState, newData.Version, newData.BackgroundUrl, newData.Id)
	}, tiktokUserEmailKey, tiktokUserIdKey, tiktokUserMobileKey, tiktokUserUsernameKey)
}

func (m *defaultUserModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *User) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "UserModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultUserModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultUserModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultUserModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*User, error) {

	builder = builder.Columns(userRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, error) {

	builder = builder.Columns(userRows)

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

	var resp []*User
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userRows)

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

	var resp []*User
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*User, error) {

	builder = builder.Columns(userRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*User, error) {

	builder = builder.Columns(userRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultUserModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultUserModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *User) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	tiktokUserEmailKey := fmt.Sprintf("%s%v", cacheTiktokUserEmailPrefix, data.Email)
	tiktokUserIdKey := fmt.Sprintf("%s%v", cacheTiktokUserIdPrefix, data.Id)
	tiktokUserMobileKey := fmt.Sprintf("%s%v", cacheTiktokUserMobilePrefix, data.Mobile)
	tiktokUserUsernameKey := fmt.Sprintf("%s%v", cacheTiktokUserUsernamePrefix, data.Username)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.Username, newData.Email, newData.Nickname, newData.Gender, newData.Role, newData.Status, newData.Mobile, newData.Password, newData.Dec, newData.Avatar, newData.DeleteTime, newData.DelState, newData.Version, newData.BackgroundUrl, newData.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, newData.Username, newData.Email, newData.Nickname, newData.Gender, newData.Role, newData.Status, newData.Mobile, newData.Password, newData.Dec, newData.Avatar, newData.DeleteTime, newData.DelState, newData.Version, newData.BackgroundUrl, newData.Id, oldVersion)
	}, tiktokUserEmailKey, tiktokUserIdKey, tiktokUserMobileKey, tiktokUserUsernameKey)
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
func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	tiktokUserEmailKey := fmt.Sprintf("%s%v", cacheTiktokUserEmailPrefix, data.Email)
	tiktokUserIdKey := fmt.Sprintf("%s%v", cacheTiktokUserIdPrefix, id)
	tiktokUserMobileKey := fmt.Sprintf("%s%v", cacheTiktokUserMobilePrefix, data.Mobile)
	tiktokUserUsernameKey := fmt.Sprintf("%s%v", cacheTiktokUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, tiktokUserEmailKey, tiktokUserIdKey, tiktokUserMobileKey, tiktokUserUsernameKey)
	return err
}
func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokUserIdPrefix, primary)
}
func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
