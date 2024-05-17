package logic

import (
	"context"
	"encoding/json"

	"tiktok/feed/rpc/feed"

	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoLogic) CreateVideo(req *types.CreateVideoReq) error {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err := l.svcCtx.FeedRpc.CreateVideo(l.ctx, &feed.CreateVideoRequest{
		ActorId:  uint32(uid),
		CoverUrl: req.CoverUrl,
		Url:      req.Url,
		Title:    req.Title,
		Category: uint32(req.Category),
	})
	if err != nil {
		return err
	}
	return nil
}
