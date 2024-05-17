package logic

import (
	"context"
	"tiktok/feed/rpc/feed"

	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideosListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideosListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideosListLogic {
	return &VideosListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideosListLogic) VideosList() (resp *types.VideosListResp, err error) {
	videos, err := l.svcCtx.FeedRpc.ListVideos(l.ctx, &feed.ListFeedRequest{ActorId: uint32(1)})
	if err != nil {
		return nil, err
	}
	resList := make([]types.VideoInfo, 0)
	for _, item := range videos.VideoList {
		resList = append(resList, types.VideoInfo{
			VideoId: int64(item.Id),
			Author: types.UserInfo{
				Id:              item.Author.Id,
				Name:            item.Author.Nickname,
				Avatar:          *item.Author.Avatar,
				Gender:          item.Author.Gender,
				Signature:       *item.Author.Signature,
				BackgroundImage: *item.Author.BackgroundImage,
				FollowCount:     *item.Author.FollowerCount,
				FollowerCount:   *item.Author.FollowCount,
				TotalFavorited:  *item.Author.TotalFavorited,
				WorkCount:       *item.Author.WorkCount,
				FavoriteCount:   *item.Author.FavoriteCount,
				IsFollow:        item.Author.IsFollow,
				FriendCount:     item.Author.FriendCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: int64(item.FavoriteCount),
			CommentCount:  int64(item.CommentCount),
			StarCount:     int64(item.StarCount),
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
			CreateTime:    item.CreateTime,
			Duration:      item.Duration,
			IsStar:        item.IsStar,
		})
	}
	return &types.VideosListResp{
		Videos: resList,
	}, nil
}
