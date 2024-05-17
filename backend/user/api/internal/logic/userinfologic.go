package logic

import (
	"context"
	"encoding/json"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId:  uid,
		ActorId: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	userInfo := types.User{
		Id:              res.User.Id,
		Nickname:        res.User.Name,
		Avatar:          *res.User.Avatar,
		Gender:          res.User.Gender,
		Signature:       *res.User.Signature,
		FollowCount:     *res.User.FollowCount,
		FollowerCount:   *res.User.FollowerCount,
		TotalFavorited:  *res.User.TotalFavorited,
		FavoriteCount:   *res.User.FavoriteCount,
		WorkCount:       *res.User.WorkCount,
		IsFollow:        res.User.IsFollow,
		BackgroundImage: *res.User.BackgroundImage,
		FriendCount:     res.User.FriendCount,
	}
	return &types.UserInfoResponse{User: userInfo}, nil
}
