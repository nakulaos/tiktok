package logic

import (
	"context"
	"encoding/json"
	"tiktok/feed/rpc/feed"

	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchEsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchEsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchEsLogic {
	return &SearchEsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchEsLogic) SearchEs(req *types.SearchEsReq) (resp *types.SearchEsResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	videos, err := l.svcCtx.FeedRpc.SearchES(l.ctx, &feed.EsSearchReq{
		UserId:  int32(uid),
		Content: req.Content,
	})

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
	return &types.SearchEsResp{
		VideoList: resList,
	}, nil
}
