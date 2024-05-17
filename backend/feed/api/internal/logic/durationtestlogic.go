package logic

import (
	"context"
	"tiktok/feed/rpc/feed"

	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DurationTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDurationTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DurationTestLogic {
	return &DurationTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DurationTestLogic) DurationTest(req *types.DurationTestReq) error {
	_, err := l.svcCtx.FeedRpc.VideoDuration(l.ctx, &feed.VideoDurationReq{
		Duration: req.Duration,
		VideoId:  uint32(req.Vid),
	})
	return err
}
