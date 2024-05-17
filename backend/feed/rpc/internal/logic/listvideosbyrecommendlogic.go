package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/common/constant"
	"tiktok/common/errorcode"
	"tiktok/common/gorse"
	"tiktok/common/utils"
	"tiktok/feed/historymodel"
	"tiktok/user/rpc/user"

	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosByRecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosByRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosByRecommendLogic {
	return &ListVideosByRecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideosByRecommendLogic) ListVideosByRecommend(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	baseurl := l.svcCtx.Config.RecommendUrl + "/api/recommend"
	url := fmt.Sprintf("%s/%d?write-back-type=read&n=%d", baseurl, in.ActorId, in.Num)
	resp, err := utils.Get(url)
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "recommend system error: %v", err)
	}

	var result []gorse.PopularResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "Json unmarshal error: %v", err)
	}

	if in.Offset == 0 {
		id, _ := strconv.Atoi(result[0].Id)
		_, err := l.svcCtx.HistoryModel.Insert(l.ctx, nil, &historymodel.History{
			Uid: int64(in.ActorId),
			Vid: int64(id),
		})
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "insert into history table failed:%v", err)
		}
	} else {
		_, err := l.svcCtx.HistoryModel.Insert(l.ctx, nil, &historymodel.History{
			Uid: int64(in.ActorId),
			Vid: int64(in.ReadVid),
		})
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "insert into history table failed:%v", err)
		}
	}

	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range result {
		id, err := strconv.Atoi(item.Id)
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "strconv error: %v", err)
		}

		video, err := l.svcCtx.VideosModel.FindOne(l.ctx, int64(id))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "find video error: %v", err)
		}

		authorInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: int64(video.AuthorId)})
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

		isFavorite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, int64(in.ActorId), int64(video.Id))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "get favorite error: %v", err)
		}

		isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, int64(in.ActorId), int64(video.Id))
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

		VideoList = append(VideoList, videoInfo)

	}

	return &feed.ListFeedResponse{}, nil
}
