package logic

import (
	"context"
	"encoding/json"
	"tiktok/common/gorse"
	"tiktok/favorite/rpc/favorite"

	"tiktok/favorite/api/internal/svc"
	"tiktok/favorite/api/internal/types"

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
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()
	if req.ActionType == 1 {
		err := gorse.Feedback(l.svcCtx.Config.RecommendUrl, "like", int(req.VideoId), int(userId))
		if err != nil {
			l.Logger.Errorf("gorse feedback like videoId:%d userId:%d err:%v", req.VideoId, userId, err)
		} else {
			l.Logger.Infof("gorse feedback like videoId:%d userId:%d", req.VideoId, userId)
		}
	}

	_, err := l.svcCtx.FavoriteRpc.FavoriteAction(l.ctx, &favorite.FavoriteActionRequest{
		UserId:     userId,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})

	return err
}
