package historymodel

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistoryModel = (*customHistoryModel)(nil)

type (
	// HistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistoryModel.
	HistoryModel interface {
		historyModel
		FindHistoryFromUid(ctx context.Context, uid int64) ([]*History, error)
	}

	customHistoryModel struct {
		*defaultHistoryModel
	}
)

func (c *customHistoryModel) FindHistoryFromUid(ctx context.Context, uid int64) ([]*History, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ?", uid)
	return c.FindAll(ctx, sb, "id DESC")
}

// NewHistoryModel returns a model for the database table.
func NewHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) HistoryModel {
	return &customHistoryModel{
		defaultHistoryModel: newHistoryModel(conn, c),
	}
}
