package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/errorcode"
	relationModel "tiktok/relation/model"
	"tiktok/user/rpc/user"

	"tiktok/relation/rpc/internal/svc"
	"tiktok/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *relation.FriendListReq) (*relation.FriendListResp, error) {
	// 互相关注的好友列表
	freinds, err := l.svcCtx.RelationModel.FindFriend(l.ctx, in.ActUser)

	if err != nil {
		if err == relationModel.ErrNotFound {
			return nil, nil
		} else {
			return nil, errors.Wrapf(errorcode.DatabaseError, "Database query user favorite list error")
		}
	}

	list := make([]*relation.UserInfo, 0)

	for _, item := range freinds {
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  in.Uid,
			ActorId: int64(item.FollowingId),
		})
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServiceUnavailable, "Get user info error from request rpc error:%s", err.Error())
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

	return &relation.FriendListResp{List: list}, nil

}
