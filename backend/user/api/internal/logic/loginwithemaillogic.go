package logic

import (
	"context"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithEmailLogic {
	return &LoginWithEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithEmailLogic) LoginWithEmail(req *types.LoginWithEmailReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
