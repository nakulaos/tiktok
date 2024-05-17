package logic

import (
	"context"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithEmailLogic {
	return &LoginWithEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithEmailLogic) LoginWithEmail(in *user.LoginWithEmailRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{}, nil
}
