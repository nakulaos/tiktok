package starmodel

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarsModel = (*customStarsModel)(nil)

type (
	// StarsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarsModel.
	StarsModel interface {
		starsModel
		GetVideoCount(ctx context.Context, uid int64) (int64, error)
		IsStar(ctx context.Context, uid, vid int64) (bool, error)
		FindIdByUidAndVid(ctx context.Context, uid, vid int64) (int64, error)
		FindOwnStars(ctx context.Context, uid int64) ([]*Stars, error)
	}

	customStarsModel struct {
		*defaultStarsModel
	}
)

func (c *customStarsModel) GetVideoCount(ctx context.Context, uid int64) (int64, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ?", uid)
	return c.FindCount(ctx, sb, "id")
}

func (c *customStarsModel) FindIdByUidAndVid(ctx context.Context, uid, vid int64) (int64, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ? AND vid = ?", uid, vid).Columns("id")
	query, values, err := sb.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = c.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (c *customStarsModel) FindOwnStars(ctx context.Context, uid int64) ([]*Stars, error) {
	sb := squirrel.Select().From(c.table)
	sb = sb.Where("uid = ?", uid)
	res, err := c.FindAll(ctx, sb, "id")
	return res, err
}

func (c *customStarsModel) IsStar(ctx context.Context, uid, vid int64) (bool, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ? AND vid = ?", uid, vid)
	cnt, err := c.FindCount(ctx, sb, "id")
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

// NewStarsModel returns a videosmodel for the database table.
func NewStarsModel(conn sqlx.SqlConn, c cache.CacheConf) StarsModel {
	return &customStarsModel{
		defaultStarsModel: newStarsModel(conn, c),
	}
}
