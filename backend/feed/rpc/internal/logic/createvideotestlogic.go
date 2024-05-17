package logic

import (
	"context"

	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoTestLogic {
	return &CreateVideoTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoTestLogic) CreateVideoTest(in *feed.CreateVideoRequest) (*feed.CreateVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &feed.CreateVideoResponse{}, nil
}
