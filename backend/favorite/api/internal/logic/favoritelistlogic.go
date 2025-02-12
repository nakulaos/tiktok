package logic

import (
	"context"
	"encoding/json"
	"tiktok/favorite/rpc/favorite"

	"tiktok/favorite/api/internal/svc"
	"tiktok/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.ListReq) (resp *types.ListResp, err error) {
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.FavoriteRpc.FavoriteList(l.ctx, &favorite.FavoriteListRequest{
		UserId:  userId,
		ActorId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	resLists := make([]types.VideoInfo, 0)
	for i := 0; i < len(res.VideoList); i++ {
		newUser := types.User{
			Id:             res.VideoList[i].Author.Id,
			Name:           res.VideoList[i].Author.Nickname,
			Gender:         res.VideoList[i].Author.Gender,
			FollowCount:    *res.VideoList[i].Author.FollowCount,
			FollowerCount:  *res.VideoList[i].Author.FollowerCount,
			IsFollow:       res.VideoList[i].Author.IsFollow,
			Avatar:         *res.VideoList[i].Author.Avatar,
			Signature:      *res.VideoList[i].Author.Signature,
			TotalFavorited: *res.VideoList[i].Author.TotalFavorited,
			WorkCount:      *res.VideoList[i].Author.WorkCount,
			FavoriteCount:  *res.VideoList[i].Author.FavoriteCount,
			FriendCount:    res.VideoList[i].Author.FriendCount,
		}
		videoDetail := types.VideoInfo{
			VideoId:       res.VideoList[i].Id,
			User:          newUser,
			PlayUrl:       res.VideoList[i].PlayUrl,
			CommentCount:  res.VideoList[i].CommentCount,
			CoverUrl:      res.VideoList[i].CoverUrl,
			FavoriteCount: res.VideoList[i].FavoriteCount,
			StarCount:     res.VideoList[i].StarCount,
			IsFavorite:    res.VideoList[i].IsFavorite,
			IsStar:        res.VideoList[i].IsStar,
			Title:         res.VideoList[i].Title,
			CreateTime:    res.VideoList[i].CreateTime,
			Duration:      res.VideoList[i].Duration,
		}

		resLists = append(resLists, videoDetail)
	}

	return &types.ListResp{VideoList: resLists}, nil
}
