package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/favorite/errorcode"
	"tiktok/favorite/rpc/favorite"
	"tiktok/favorite/rpc/internal/svc"
	"tiktok/favorite/starmodel"
	videosModel "tiktok/feed/videosmodel"
	userModel "tiktok/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStarActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarActionLogic {
	return &StarActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StarActionLogic) StarAction(in *favorite.StarActionRequest) (*favorite.StarActionResponse, error) {
	userID := in.UserId
	videoID := in.VideoId
	actionType := in.ActionType

	// 1. 判断用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.FavoriteUserIdEmptyError, "user is not found , id = %d", userID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user failed, err = %v", err)
		}
	}

	// 2.检测视频是否存在
	_, err = l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.FavoriteVideoIdEmptyError, "video is not found , id = %d", videoID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find video failed, err = %v", err)
		}
	}

	switch actionType {
	case 1:
		isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, userID, videoID)
		if err != nil {
			if err == starmodel.ErrNotFound {
				logx.WithContext(l.ctx).Infof("the record is not found in the database,user_id=%d,video_id=%d", userID, videoID)
			} else {
				return nil, errors.Wrapf(errorcode2.ServerError, "find favorite record failed,err=%v", err)
			}
		}

		if isStar {
			return nil, errors.Wrapf(errorcode.StarServiceDuplicateError, "the record is already exist,user_id=%d,video_id=%d", userID, videoID)
		}

		newStar := starmodel.Stars{
			Uid: userID,
			Vid: videoID,
		}

		err = l.svcCtx.VideosModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			videoDetail, err := l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
			if err != nil {
				return err
			}

			_, err = l.svcCtx.StarModel.Insert(l.ctx, nil, &newStar)
			if err != nil {
				return err
			}

			_, err = l.svcCtx.VideosModel.Update(l.ctx, nil, &videosModel.Videos{
				Id:            videoID,
				AuthorId:      videoDetail.AuthorId,
				Title:         videoDetail.Title,
				CoverUrl:      videoDetail.CoverUrl,
				PlayUrl:       videoDetail.Title,
				FavoriteCount: videoDetail.FavoriteCount,
				StarCount:     videoDetail.StarCount + 1,
				CommentCount:  videoDetail.CommentCount,
				DelState:      videoDetail.DelState,
				Category:      videoDetail.Category,
				Version:       videoDetail.Version,
				Duration:      videoDetail.Duration,
				DeleteTime:    videoDetail.DeleteTime,
			})

			return err

		})

		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "[favorite] Transaction error:%v", err)
		}

	case 2:
		isStar, err := l.svcCtx.StarModel.IsStar(l.ctx, userID, videoID)
		if err != nil {
			if err == starmodel.ErrNotFound {
				return nil, errors.Wrapf(errorcode.StarServiceCancelError, "the star record is not exist,user_id=%d,video_id=%d", userID, videoID)
			} else {
				return nil, errors.Wrapf(errorcode2.ServerError, "find favorite record failed,err=%v", err)
			}
		}

		if !isStar {
			return nil, errors.Wrapf(errorcode.StarServiceCancelError, "the star record is not exist,user_id=%d,video_id=%d", userID, videoID)
		}

		err = l.svcCtx.StarModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			videoDetail, err := l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
			if err != nil {
				return err
			}
			id, err := l.svcCtx.StarModel.FindIdByUidAndVid(l.ctx, userID, videoID)
			if err != nil {
				return err
			}
			err = l.svcCtx.StarModel.Delete(l.ctx, nil, id)
			if err != nil {
				return err
			}

			_, err = l.svcCtx.VideosModel.Update(l.ctx, nil, &videosModel.Videos{
				Id:            videoID,
				AuthorId:      videoDetail.AuthorId,
				Title:         videoDetail.Title,
				CoverUrl:      videoDetail.CoverUrl,
				PlayUrl:       videoDetail.Title,
				FavoriteCount: videoDetail.FavoriteCount,
				StarCount:     videoDetail.StarCount - 1,
				CommentCount:  videoDetail.CommentCount,
				DelState:      videoDetail.DelState,
				Category:      videoDetail.Category,
				Version:       videoDetail.Version,
				Duration:      videoDetail.Duration,
				DeleteTime:    videoDetail.DeleteTime,
			})

			return err
		})

		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "[favorite] Transaction error:%v", err)
		}

	}

	return &favorite.StarActionResponse{}, nil
}
