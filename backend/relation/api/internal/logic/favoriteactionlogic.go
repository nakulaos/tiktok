package logic

import (
	"context"
	"encoding/json"
	"tiktok/relation/rpc/relation"

	"tiktok/relation/api/internal/svc"
	"tiktok/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.ActionReq) error {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	_, err := l.svcCtx.RelationRpc.Favorite(l.ctx, &relation.FavoriteRequest{
		Uid:    uid,
		ToUid:  req.ToUserID,
		Action: req.Action,
	})

	return err
}
