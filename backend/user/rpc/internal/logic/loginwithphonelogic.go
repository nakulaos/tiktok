package logic

import (
	"context"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithPhoneLogic {
	return &LoginWithPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginWithPhoneLogic) LoginWithPhone(in *user.LoginWithPhoneRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{}, nil
}
