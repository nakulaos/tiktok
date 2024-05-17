package favoritemodel

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoritesModel = (*customFavoritesModel)(nil)

type (
	// FavoritesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoritesModel.
	FavoritesModel interface {
		favoritesModel
		GetVideoCount(ctx context.Context, uid int64) (int64, error)
		IsFavorite(ctx context.Context, uid, vid int64) (bool, error)
		FindIdByUidAndVid(ctx context.Context, uid, vid int64) (int64, error)
		FindOwnFavorites(ctx context.Context, uid int64) ([]*Favorites, error)
	}

	customFavoritesModel struct {
		*defaultFavoritesModel
	}
)

func (c *customFavoritesModel) FindOwnFavorites(ctx context.Context, uid int64) ([]*Favorites, error) {
	sb := squirrel.Select().From(c.table)
	sb = sb.Where("uid = ?", uid)
	res, err := c.FindAll(ctx, sb, "id")
	return res, err
}

func (c *customFavoritesModel) FindIdByUidAndVid(ctx context.Context, uid, vid int64) (int64, error) {
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

func (c *customFavoritesModel) IsFavorite(ctx context.Context, uid, vid int64) (bool, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ? AND vid = ?", uid, vid)
	cnt, err := c.FindCount(ctx, sb, "id")
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (c *customFavoritesModel) GetVideoCount(ctx context.Context, uid int64) (int64, error) {
	sb := squirrel.Select().From(c.table).Where("uid = ?", uid)
	return c.FindCount(ctx, sb, "id")
}

// NewFavoritesModel returns a favoritemodel for the database table.
func NewFavoritesModel(conn sqlx.SqlConn, c cache.CacheConf) FavoritesModel {
	return &customFavoritesModel{
		defaultFavoritesModel: newFavoritesModel(conn, c),
	}
}
