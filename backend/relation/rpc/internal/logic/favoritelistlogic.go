package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"tiktok/common/errorcode"
	relationModel "tiktok/relation/model"
	"tiktok/user/rpc/user"

	"tiktok/relation/rpc/internal/svc"
	"tiktok/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *relation.FavoriteListReq) (*relation.FavoriteListResp, error) {
	// 用户关注列表的用户信息
	build := squirrel.Select().From("`relations`").Where("follower_id = ?", in.ActUser)
	favorite, err := l.svcCtx.RelationModel.FindAll(l.ctx, build, "")
	if err != nil {
		if err == relationModel.ErrNotFound {
			return nil, nil
		} else {
			return nil, errors.Wrapf(errorcode.DatabaseError, "Database query user favorite list error")
		}
	}

	list := make([]*relation.UserInfo, 0)

	for _, item := range favorite {
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  in.Uid,
			ActorId: int64(item.FollowingId),
		})
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServiceUnavailable, "Get user info error from request rpc error: %+v", err)
		}

		var coverUri string
		var vid int64
		video, err := l.svcCtx.VideosModel.FindLastByUid(l.ctx, int64(userInfo.User.Id))
		if err != nil {
			coverUri = ""
			vid = 0
		} else {
			coverUri = video.CoverUrl
			vid = video.Id
		}

		list = append(list, &relation.UserInfo{
			Id:              int64(userInfo.User.Id),
			NickName:        userInfo.User.Name,
			Gender:          int64(userInfo.User.Gender),
			Avatar:          *userInfo.User.Avatar,
			Dec:             *userInfo.User.Signature,
			BackgroundImage: *userInfo.User.BackgroundImage,
			VideoId:         vid,
			CoverUrl:        coverUri,
			FollowCount:     *userInfo.User.FollowCount,
			FollowerCount:   *userInfo.User.FollowerCount,
			IsFollow:        userInfo.User.IsFollow,
			TotalFavorited:  *userInfo.User.TotalFavorited,
			WorkCount:       *userInfo.User.WorkCount,
			FavoriteCount:   *userInfo.User.FavoriteCount,
			FriendCount:     userInfo.User.FriendCount,
		})

	}

	return &relation.FavoriteListResp{List: list}, nil
}
