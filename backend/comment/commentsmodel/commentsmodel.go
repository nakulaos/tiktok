package commentsmodel

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
	"tiktok/common/globalkey"
	"time"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
		InsertWithSqlTx(context context.Context, tx *sql.Tx, data *Comments) (int64, error)
		DeleteWithSqlTx(context context.Context, tx *sql.Tx, data *Comments) (int64, error)
		IsCommentExist(ctx context.Context, id int64) (bool, error)
		FindComment(ctx context.Context, uid, vid int64, content string) (int64, error)
		FindCommentsByVid(ctx context.Context, vid int64) ([]*Comments, error)
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}
)

func (c *customCommentsModel) FindCommentsByVid(ctx context.Context, vid int64) ([]*Comments, error) {
	sb := squirrel.Select().From(c.table).Where("vid = ? and del_state = ?", vid, globalkey.DelStateNo)
	return c.FindAll(ctx, sb, "")

}

func (c *customCommentsModel) FindComment(ctx context.Context, uid, vid int64, content string) (int64, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ? and vid = ? and content = ?", uid, vid, content)
	ress, err := c.FindAll(ctx, sb, "")
	if err != nil {
		return 0, err
	}
	if len(ress) > 0 {
		return ress[0].Id, nil
	}

	return 0, nil

}

func (c *customCommentsModel) IsCommentExist(ctx context.Context, id int64) (bool, error) {
	sb := squirrel.Select().From(c.table).Where("id = ?", id)
	cnt, err := c.FindCount(ctx, sb, "id")
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (c *customCommentsModel) DeleteWithSqlTx(context context.Context, tx *sql.Tx, data *Comments) (int64, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ? and vid = ? and del_state = ? and content = ?", data.Uid, data.Vid, globalkey.DelStateNo, data.Content)
	comment, err := c.FindAll(context, sb, "")
	if err == sqlc.ErrNotFound {
		return 0, nil
	}
	query := fmt.Sprintf("delete from %s where `id` = ?", c.table)

	_, err = tx.ExecContext(context, query, comment[0].Id)
	// 删除对应缓存
	tiktokCommentsIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentsIdPrefix, comment[0].Id)
	_ = c.DelCacheCtx(context, tiktokCommentsIdKey)
	return comment[0].Id, err
}

func (c *customCommentsModel) InsertWithSqlTx(context context.Context, tx *sql.Tx, data *Comments) (int64, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", c.table, commentsRowsExpectAutoSet)
	res, err := tx.ExecContext(context, query, data.Uid, data.Vid, data.Content, data.DeleteTime, data.DelState, data.Version)
	// 删除对应缓存
	id, _ := res.LastInsertId()
	tiktokCommentsIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentsIdPrefix, id)
	_ = c.DelCacheCtx(context, tiktokCommentsIdKey)

	return id, err
}

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn sqlx.SqlConn, c cache.CacheConf) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn, c),
	}
}
