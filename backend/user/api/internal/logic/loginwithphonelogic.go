package logic

import (
	"context"

	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithPhoneLogic {
	return &LoginWithPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithPhoneLogic) LoginWithPhone(req *types.LoginWithPhoneReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
