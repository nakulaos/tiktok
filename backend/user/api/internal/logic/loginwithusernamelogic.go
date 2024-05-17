package logic

import (
	"context"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithUsernameLogic {
	return &LoginWithUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithUsernameLogic) LoginWithUsername(req *types.LoginWithUsernameReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.LoginWithUsername(l.ctx, &user.LoginWithUsernameRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{
		Avatar:          res.Avatar,
		AccessToken:     res.Token,
		UserID:          res.UserId,
		Name:            res.Nickname,
		Gender:          uint32(res.Gender),
		Signature:       res.Signature,
		Username:        req.Username,
		BackgroundImage: res.BackgroundImage,
	}

	return
}
