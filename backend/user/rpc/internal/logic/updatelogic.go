package logic

import (
	"context"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户信息
func (l *UpdateLogic) Update(in *user.UpdateReq) (*user.UpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateResp{}, nil
}
