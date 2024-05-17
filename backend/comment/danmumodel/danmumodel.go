package danmumodel

import (
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
)

var _ DanmuModel = (*customDanmuModel)(nil)

type (
	// DanmuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDanmuModel.
	DanmuModel interface {
		danmuModel
		FindDanmusByVid(ctx context.Context, vid int64) ([]*Danmu, error)
	}

	customDanmuModel struct {
		*defaultDanmuModel
	}
)

func (c *customDanmuModel) FindDanmusByVid(ctx context.Context, vid int64) ([]*Danmu, error) {
	sb := squirrel.Select().From(c.table).Where("vid=?", vid)
	return c.FindAll(ctx, sb, "")
}

// NewDanmuModel returns a model for the database table.
func NewDanmuModel(conn sqlx.SqlConn, c cache.CacheConf) DanmuModel {
	return &customDanmuModel{
		defaultDanmuModel: newDanmuModel(conn, c),
	}
}
