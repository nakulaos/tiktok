package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationsModel = (*customRelationsModel)(nil)

type (
	// RelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationsModel.
	RelationsModel interface {
		relationsModel
		FindFriend(ctx context.Context, uid int64) ([]*Relations, error)
		FindRelation(ctx context.Context, userID int64, actionID int64) (int64, error)
		IsFollow(ctx context.Context, userId int64, actionId int64) (bool, error)
	}

	customRelationsModel struct {
		*defaultRelationsModel
	}
)

func (c *customRelationsModel) FindFriend(ctx context.Context, uid int64) ([]*Relations, error) {
	// follower_id 关注人
	// following_id 被关注人
	var friends []*Relations
	query := fmt.Sprintf("SELECT %s FROM %s WHERE follower_id  IN  (SELECT following_id FROM %s WHERE follower_id = ?)AND following_id = ?", relationsRows, c.table, c.table)
	err := c.conn.QueryRowsPartialCtx(ctx, &friends, query, uid, uid)
	switch err {
	case nil:
		return friends, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c *customRelationsModel) FindRelation(ctx context.Context, userID int64, actionID int64) (int64, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE follower_id = ? AND following_id = ?", c.table)
	var id int64
	err := c.conn.QueryRowPartialCtx(ctx, &id, query, userID, actionID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *customRelationsModel) IsFollow(ctx context.Context, userId int64, actionId int64) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE follower_id = ? AND following_id = ?", c.table)
	var count int64
	err := c.conn.QueryRowPartialCtx(ctx, &count, query, userId, actionId)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// NewRelationsModel returns a favoritemodel for the database table.
func NewRelationsModel(conn sqlx.SqlConn) RelationsModel {
	return &customRelationsModel{
		defaultRelationsModel: newRelationsModel(conn),
	}
}
