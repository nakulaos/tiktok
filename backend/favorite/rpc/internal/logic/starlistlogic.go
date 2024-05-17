package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/constant"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/favorite/errorcode"
	videosModel "tiktok/feed/videosmodel"
	userModel "tiktok/user/model"
	"tiktok/user/rpc/user"

	"tiktok/favorite/rpc/favorite"
	"tiktok/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarListLogic {
	return &StarListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StarListLogic) StarList(in *favorite.StarListRequest) (*favorite.StarListResponse, error) {
	userId := in.UserId
	actorId := in.ActorId

	if _, err := l.svcCtx.UserModel.FindOne(l.ctx, userId); err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.StarUserIdEmptyError, "the user is not exist , id : %d", userId)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user error, id : %d,err :%s", in.UserId, err.Error())
		}
	}

	if _, err := l.svcCtx.UserModel.FindOne(l.ctx, actorId); err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.StarUserIdEmptyError, "the user is not exist , id : %d", actorId)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user error, id : %d,err :%s", in.ActorId, err.Error())
		}
	}

	favorVideos, err := l.svcCtx.StarModel.FindOwnStars(l.ctx, actorId)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "find favorite error, actorId : %d,err :%s", actorId, err.Error())
	}

	videoInfoList := make([]*favorite.Video, 0)

	for i := 0; i < len(favorVideos); i++ {
		videoid := favorVideos[i].Vid
		videoDetail, err := l.svcCtx.VideosModel.FindOne(l.ctx, videoid)
		if err != nil {
			if err == videosModel.ErrNotFound {
				return nil, errors.Wrapf(errorcode.StarVideoIdEmptyError, "video is not exist, id : %d", videoid)
			} else {
				return nil, errors.Wrapf(errorcode2.ServerError, "find video error, id : %d,err :%s", videoid, err.Error())
			}
		}

		// 获取视频作者信息
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: userId, ActorId: videoDetail.AuthorId})
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "get user info error, id : %d,err :%s", videoDetail.AuthorId, err.Error())
		}

		isFavorited, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, userId, videoid)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "check favorite error, id : %d,err :%s", videoid, err.Error())
		}

		isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, userId, videoid)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "check star error, id : %d,err :%s", videoid, err.Error())
		}

		userDetail := &favorite.User{
			Id:             userInfo.User.Id,
			Nickname:       userInfo.User.Name,
			Gender:         userInfo.User.Gender,
			FollowCount:    userInfo.User.FollowCount,
			FollowerCount:  userInfo.User.FollowerCount,
			IsFollow:       userInfo.User.IsFollow,
			Avatar:         userInfo.User.Avatar,
			Signature:      userInfo.User.Signature,
			TotalFavorited: userInfo.User.TotalFavorited,
			WorkCount:      userInfo.User.WorkCount,
			FavoriteCount:  userInfo.User.FavoriteCount,
			FriendCount:    userInfo.User.FriendCount,
		}

		videoInfo := &favorite.Video{
			Id:            int64(videoDetail.Id),
			Author:        userDetail,
			PlayUrl:       videoDetail.PlayUrl,
			CoverUrl:      videoDetail.CoverUrl,
			FavoriteCount: int64(videoDetail.FavoriteCount),
			CommentCount:  int64(videoDetail.CommentCount),
			StarCount:     int64(videoDetail.StarCount),
			IsFavorite:    isFavorited,
			IsStar:        isStar,
			Title:         videoDetail.Title,
			CreateTime:    videoDetail.CreateTime.Format(constant.TimeFormat),
			Duration:      videoDetail.Duration.String,
		}

		videoInfoList = append(videoInfoList, videoInfo)
	}

	return &favorite.StarListResponse{VideoList: videoInfoList}, nil
}
