package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/constant"
	"tiktok/common/errorcode"
	"tiktok/feed/historymodel"
	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"
	"tiktok/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVideoLogic {
	return &FindVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindVideoLogic) FindVideo(in *feed.FindVideoReq) (*feed.FindVideoResp, error) {

	video, err := l.svcCtx.VideosModel.FindOne(l.ctx, int64(in.Vid))
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "find video error: %v", err)
	}

	// 获取该视频的作者信息
	authorInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId:  int64(in.Uid),
		ActorId: int64(video.AuthorId),
	})

	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "get user info error: %v", err)
	}

	isFavorite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, int64(in.Uid), int64(in.Vid))
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "get favorite error: %v", err)
	}

	isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, int64(in.Uid), int64(in.Vid))
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "get star error: %v", err)
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

	videoInfo := &feed.VideoInfo{
		Id:            uint32(in.Vid),
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

	// 插入历史记录
	_, err = l.svcCtx.HistoryModel.Insert(l.ctx, nil, &historymodel.History{
		Uid: int64(in.Uid),
		Vid: int64(in.Vid),
	})

	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "insert history error: %v", err)
	}

	return &feed.FindVideoResp{Video: videoInfo}, nil
}
