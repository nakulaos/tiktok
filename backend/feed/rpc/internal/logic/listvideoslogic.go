package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/constant"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/feed/errorcode"
	videosModel "tiktok/feed/videosmodel"
	"tiktok/user/rpc/user"

	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosLogic {
	return &ListVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideosLogic) ListVideos(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	// 找最新的视频
	videos, err := l.svcCtx.VideosModel.FindNewVideos(l.ctx)
	if err != nil {
		if err == videosModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.FeedUnableToQueryVideoError, "the video is not exist.err:%v,err:%v", err)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "query in database failed.err:%v", err)
		}
	}

	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range videos {
		userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: int64(item.AuthorId)})
		if err != nil {
			if e, ok := err.(errorcode2.ErrorCode); ok {
				return nil, e
			} else {
				return nil, errors.Wrapf(errorcode2.ServerError, "failed to get user info.err:%v", err)
			}

		}
		userInfo := &feed.User{
			Id:              userRpcRes.User.Id,
			Nickname:        userRpcRes.User.Name,
			FollowCount:     userRpcRes.User.FollowCount,
			FollowerCount:   userRpcRes.User.FollowCount,
			IsFollow:        false,
			Avatar:          userRpcRes.User.Avatar,
			BackgroundImage: userRpcRes.User.BackgroundImage,
			Signature:       userRpcRes.User.Signature,
			TotalFavorited:  userRpcRes.User.TotalFavorited,
			WorkCount:       userRpcRes.User.WorkCount,
			FavoriteCount:   userRpcRes.User.FavoriteCount,
			Gender:          userRpcRes.User.Gender,
			FriendCount:     userRpcRes.User.FriendCount,
		}

		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(item.Id),
			Author:        userInfo,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: uint32(item.FavoriteCount),
			CommentCount:  uint32(item.CommentCount),
			StarCount:     uint32(item.StarCount),
			IsFavorite:    false,
			IsStar:        false,
			Title:         item.Title,
			CreateTime:    item.CreateTime.Format(constant.TimeFormat),
			Duration:      item.Duration.String,
		})
	}

	return &feed.ListFeedResponse{VideoList: VideoList}, nil
}
