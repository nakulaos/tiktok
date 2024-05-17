package logic

import (
	"context"
	"tiktok/user/rpc/user"

	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		UserName:        req.Username,
		Gender:          req.Gender,
		Phone:           &req.Phone,
		Password:        req.Password,
		Avatar:          req.Avatar,
		Dec:             req.Dec,
		BackgroundImage: req.BackgroundImage,
		Email:           &req.Email,
	})
	return err
}
