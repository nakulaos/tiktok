package logic

import (
	"context"

	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoTestLogic {
	return &CreateVideoTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoTestLogic) CreateVideoTest(req *types.CreateVideoReq) error {
	// todo: add your logic here and delete this line

	return nil
}
