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

type ListHistoryVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHistoryVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHistoryVideosLogic {
	return &ListHistoryVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHistoryVideosLogic) ListHistoryVideos(in *feed.HistoryReq) (*feed.HistoryResp, error) {
	histories, err := l.svcCtx.HistoryModel.FindHistoryFromUid(l.ctx, int64(in.Uid))
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "find videos by category failed:%v", err)
	}

	videoList := make([]*feed.VideoInfo, 0)
	for _, history := range histories {
		videoinfo, err := l.svcCtx.VideosModel.FindOne(l.ctx, int64(history.Vid))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "find video failed:%v", err)
		}

		authorInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.Uid), ActorId: int64(videoinfo.AuthorId)})
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

		isFavorite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, int64(in.Uid), int64(history.Vid))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "get favorite error: %v", err)
		}

		isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, int64(in.Uid), int64(history.Vid))
		if err != nil {
			return nil, errors.Wrapf(errorcode.ServerError, "get star error: %v", err)
		}

		videoList = append(videoList, &feed.VideoInfo{
			Id:            uint32(videoinfo.Id),
			Author:        userInfo,
			PlayUrl:       videoinfo.PlayUrl,
			CoverUrl:      videoinfo.CoverUrl,
			FavoriteCount: uint32(videoinfo.FavoriteCount),
			CommentCount:  uint32(videoinfo.CommentCount),
			StarCount:     uint32(videoinfo.StarCount),
			IsFavorite:    isFavorite,
			IsStar:        isStar,
			Title:         videoinfo.Title,
			CreateTime:    videoinfo.CreateTime.Format(constant.TimeFormat),
			Duration:      videoinfo.Duration.String,
		})
	}

	return &feed.HistoryResp{VideoList: videoList}, nil
}
