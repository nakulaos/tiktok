package videosmodel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideosModel = (*customVideosModel)(nil)

type (
	// VideosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideosModel.
	VideosModel interface {
		videosModel
		FindLastByUid(ctx context.Context, uid int64) (*Videos, error)
		GetFavariteCount(ctx context.Context, id int64) (int64, error)
		GetVideosListFromUid(ctx context.Context, uid int64) ([]*Videos, error)
		FindVideosFromCategory(ctx context.Context, category int64) ([]*Videos, error)
		FindNewVideos(ctx context.Context) ([]*Videos, error)
		IncrCommentCount(ctx context.Context, tx *sql.Tx, id int64) error
		DecrCommentCount(ctx context.Context, tx *sql.Tx, id int64) error
	}

	customVideosModel struct {
		*defaultVideosModel
	}
)

func (c *customVideosModel) IncrCommentCount(ctx context.Context, tx *sql.Tx, id int64) error {
	res, err := c.FindOne(ctx, id)
	if err != nil {
		return err
	}
	res.CommentCount++
	tiktokVideosIdKey := fmt.Sprintf("%s%v", cacheTiktokVideosIdPrefix, id)
	query := fmt.Sprintf("update %s set %s where `id` = ?", c.table, videosRowsWithPlaceHolder)
	_, err = tx.ExecContext(ctx, query, res.AuthorId, res.Title, res.CoverUrl, res.PlayUrl, res.FavoriteCount, res.StarCount, res.CommentCount, res.DeleteTime, res.DelState, res.Category, res.Version, res.Duration, res.Id)
	_ = c.DelCacheCtx(ctx, tiktokVideosIdKey)
	return err
}

func (c *customVideosModel) DecrCommentCount(ctx context.Context, tx *sql.Tx, id int64) error {
	res, err := c.FindOne(ctx, id)
	if err != nil {
		return err
	}
	res.CommentCount--
	tiktokVideosIdKey := fmt.Sprintf("%s%v", cacheTiktokVideosIdPrefix, id)
	query := fmt.Sprintf("update %s set %s where `id` = ?", c.table, videosRowsWithPlaceHolder)
	_, err = tx.ExecContext(ctx, query, res.AuthorId, res.Title, res.CoverUrl, res.PlayUrl, res.FavoriteCount, res.StarCount, res.CommentCount, res.DeleteTime, res.DelState, res.Category, res.Version, res.Duration, res.Id)
	_ = c.DelCacheCtx(ctx, tiktokVideosIdKey)
	return err
}

func (c *customVideosModel) FindNewVideos(ctx context.Context) ([]*Videos, error) {
	sb := squirrel.Select().From(c.table)
	return c.FindPageListByIdDESC(ctx, sb, 0, 15)
}

func (c *customVideosModel) FindVideosFromCategory(ctx context.Context, category int64) ([]*Videos, error) {
	sb := squirrel.Select().From(c.table).Where("category = ?", category)
	return c.FindAll(ctx, sb, "")
}

func (c *customVideosModel) GetVideosListFromUid(ctx context.Context, uid int64) ([]*Videos, error) {
	sb := squirrel.Select().From(c.table).Where("author_id = ?", uid)
	return c.FindAll(ctx, sb, "id")
}

func (c *customVideosModel) GetFavariteCount(ctx context.Context, uid int64) (int64, error) {
	var totalCount int64
	query := fmt.Sprintf("SELECT COALESCE(SUM('favorite_count'),0) FROM %s WHERE author_id = ?", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &totalCount, query, uid)
	switch err {
	case nil:
		return totalCount, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return totalCount, err
	}
}

func (c *customVideosModel) FindLastByUid(ctx context.Context, uid int64) (*Videos, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE author_id  = ? ORDER BY id DESC LIMIT 1", videosRows, c.table)
	var resp Videos
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewVideosModel returns a favoritemodel for the database table.
func NewVideosModel(conn sqlx.SqlConn, c cache.CacheConf) VideosModel {
	return &customVideosModel{
		defaultVideosModel: newVideosModel(conn, c),
	}
}
