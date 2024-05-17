package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/user/errorcode"
	userModel "tiktok/user/model"

	"tiktok/user/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户信息
func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userID := in.UserId
	actionID := in.ActorId

	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errorcode.UserNotExistError
		} else {
			return nil, errors.Wrapf(errorcode2.DatabaseError, "err:%v", err)
		}
	}

	res, err := l.svcCtx.UserModel.FindOne(l.ctx, actionID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errorcode.UserNotExistError
		} else {
			return nil, errors.Wrapf(errorcode2.DatabaseError, "err:%v", err)
		}
	}

	// TODO: 补充各种查询逻辑
	// 获取用户关注数
	sb := squirrel.Select().From("`relations`").Where("follower_id = ?", actionID)
	followCount, err := l.svcCtx.RelationsModel.FindCount(l.ctx, sb, "id")
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "Error in obtaining the number of followers,err:%+v", err)
	}
	uint32FollowCount := uint32(followCount)
	//logx.WithContext(l.ctx).Infof("uint32FollowCount : %d", uint32FollowCount)

	// 获取用户粉丝数
	sb = squirrel.Select().From("`relations`").Where("following_id = ?", actionID)
	followerCount, err := l.svcCtx.RelationsModel.FindCount(l.ctx, sb, "id")
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "Error in obtaining the number of followings")
	}
	uint32FollowerCount := uint32(followerCount)

	// 获取视频总点赞数
	favorcount, err := l.svcCtx.VideosModel.GetFavariteCount(l.ctx, actionID)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "Error in obtaining the number of favorited")
	}
	uint32favoriteCount := uint32(favorcount)

	// 获取用户视频总数
	videoCount, err := l.svcCtx.FavorModel.GetVideoCount(l.ctx, actionID)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "Error in obtaining the number of videos")
	}

	// 查询是否关注
	var isFollow bool
	if userID != actionID {
		isFollow, err = l.svcCtx.RelationsModel.IsFollow(l.ctx, userID, actionID)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "Error in obtaining the number of videos")
		}
	}

	// 查询用户所有视频
	videoLists, err := l.svcCtx.VideosModel.GetVideosListFromUid(l.ctx, actionID)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "Error in obtaining  Lists of videos,err : %s", err.Error())
	}
	workCount := uint32(len(videoLists))

	uint32favoriteVideoCount := uint32(videoCount)

	backgroundUri, _ := res.BackgroundUrl.Value()
	backgroundUriStr := backgroundUri.(string)
	userInfo := &user.UserInfo{
		Id:              uint32(res.Id),
		Name:            res.Nickname,
		FollowCount:     &uint32FollowCount,
		FollowerCount:   &uint32FollowerCount,
		IsFollow:        isFollow,
		Avatar:          &res.Avatar,
		BackgroundImage: &backgroundUriStr,
		Signature:       &res.Dec,
		TotalFavorited:  &uint32favoriteCount,
		WorkCount:       &workCount,
		FavoriteCount:   &uint32favoriteVideoCount,
		Gender:          uint32(res.Gender),
		FriendCount:     0,
	}
	//logx.WithContext(l.ctx).Infof("userinfo count:%d", *userInfo.FollowCount)
	return &user.UserInfoResponse{User: userInfo}, nil
}
