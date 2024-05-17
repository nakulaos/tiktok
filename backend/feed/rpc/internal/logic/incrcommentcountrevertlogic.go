package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiktok/common/errorcode"

	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncrCommentCountRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrCommentCountRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrCommentCountRevertLogic {
	return &IncrCommentCountRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncrCommentCountRevertLogic) IncrCommentCountRevert(in *feed.IncrCommentCountReq) (*feed.IncrCommentCountResp, error) {
	id := in.Id
	l.Logger.Infof("decretment comment count  for video id: %d", id)
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DataSource).RawDB()
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "db error:%v", err)
	}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		err = l.svcCtx.VideosModel.DecrCommentCount(l.ctx, tx, id)
		return err
	}); err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "decretment comment count failed: %v", err)
	}

	return nil, nil
}
