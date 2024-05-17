package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/constant"
	"tiktok/common/errorcode"
	"tiktok/user/rpc/user"

	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoLogic {
	return &ListVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideoLogic) ListVideo(in *feed.ListVideoRequest) (*feed.ListVideoResponse, error) {
	// 列出to-uid的所有视频
	videos, err := l.svcCtx.VideosModel.GetVideosListFromUid(l.ctx, int64(in.Uid))
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "find videos by category failed:%v", err)
	}

	videoList := make([]*feed.VideoInfo, 0)
	authorInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.Uid), ActorId: int64(in.ToUid)})
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "get user info failed:%v", err)
	}
	userInfo := &feed.User{
		Id:              authorInfo.User.Id,
		Nickname:        authorInfo.User.Name,
		FollowCount:     authorInfo.User.FollowCount,
		FollowerCount:   authorInfo.User.FollowCount,
		IsFollow:        authorInfo.User.IsFollow,
		Avatar:          authorInfo.User.Avatar,
		BackgroundImage: authorInfo.User.BackgroundImage,
		Signature:       authorInfo.User.Signature,
		TotalFavorited:  authorInfo.User.TotalFavorited,
		WorkCount:       authorInfo.User.WorkCount,
		FavoriteCount:   authorInfo.User.FavoriteCount,
		Gender:          authorInfo.User.Gender,
		FriendCount:     authorInfo.User.FriendCount,
	}
	for _, video := range videos {
		isFavorite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, int64(in.Uid), int64(video.Id))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "get favorite error: %v", err)
		}

		isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, int64(in.Uid), int64(video.Id))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "get star error: %v", err)
		}

		videoInfo := &feed.VideoInfo{
			Id:            uint32(video.Id),
			Author:        userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: uint32(video.FavoriteCount),
			CommentCount:  uint32(video.CommentCount),
			IsFavorite:    isFavorite,
			Title:         video.Title,
			CreateTime:    video.CreateTime.Format(constant.TimeFormat),
			StarCount:     uint32(video.StarCount),
			IsStar:        isStar,
			Duration:      video.Duration.String,
		}

		videoList = append(videoList, videoInfo)

	}

	return &feed.ListVideoResponse{VideoList: videoList}, nil
}
