package logic

import (
	"context"
	"encoding/json"
	"tiktok/live/rpc/live"

	"tiktok/live/api/internal/svc"
	"tiktok/live/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLogic {
	return &StartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartLogic) Start() (resp *types.StartResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	startLive, err := l.svcCtx.LiveRpc.StartLive(l.ctx, &live.StartLiveRequest{Uid: uid})
	if err != nil {
		return nil, err
	}

	return &types.StartResp{
		StreamUrl: startLive.StreamUrl,
	}, nil
}
