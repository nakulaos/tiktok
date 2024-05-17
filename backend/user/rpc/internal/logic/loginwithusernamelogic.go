package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/crypt"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/common/jwtx"
	"tiktok/user/errorcode"
	"tiktok/user/model"
	"time"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithUsernameLogic {
	return &LoginWithUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithUsernameLogic) LoginWithUsername(in *user.LoginWithUsernameRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(errorcode.UserNotExistError, "username:%s,err:username not found", in.Username)
		}

		return nil, errors.Wrapf(errorcode2.ServerError, "err:%v", err)
	}

	password := crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if res.Password != password {
		return nil, errors.Wrapf(errorcode.UserNotExistError, "req:%s", in.Username)
	}

	token, err := jwtx.GetToken(l.svcCtx.Config.JWTAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JWTAuth.AccessExpire, res.Id)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "generate token faild,username:%s", res.Username)
	}

	return &user.LoginResponse{
		UserId:          res.Id,
		Avatar:          res.Avatar,
		Nickname:        res.Nickname,
		Gender:          res.Gender,
		BackgroundImage: res.BackgroundUrl.String,
		Signature:       res.Dec,
		Token:           token,
	}, nil
}
