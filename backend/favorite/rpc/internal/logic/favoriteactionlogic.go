package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	errorcode2 "tiktok/common/errorcode"
	"tiktok/favorite/errorcode"
	favoriteModel "tiktok/favorite/favoritemodel"
	videosModel "tiktok/feed/videosmodel"
	userModel "tiktok/user/model"
	"time"

	"tiktok/favorite/rpc/favorite"
	"tiktok/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
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

	value := strconv.FormatInt(userID, 10) + "#" + strconv.FormatInt(videoID, 10)

	switch actionType {
	case 1:
		// 点赞操作
		// 利用redis的set集合，存储用户点赞

		// 1. 利用判断是否已经点赞
		exist, err := l.svcCtx.BizRedis.SismemberCtx(l.ctx, "tiktok:favorite:like:", value)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "redis sismember failed in FavoriteAction function, err = %v", err)
		}

		if exist {
			// 已经存在
			return nil, errors.Wrapf(errorcode.FavoriteServiceDuplicateError, "user %d has liked video %d", userID, videoID)
		} else {
			// 不存在
			// 去查数据库
			isFavourite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, userID, videoID)
			if err != nil {
				if err == favoriteModel.ErrNotFound {
					logx.WithContext(l.ctx).Infof("the record is not found in the database,user_id=%d,video_id=%d", userID, videoID)
				} else {
					return nil, errors.Wrapf(errorcode2.ServerError, "find favorite record failed,err=%v", err)
				}
			}
			if isFavourite {
				// 写入缓存
				_, err := l.svcCtx.BizRedis.SaddCtx(l.ctx, "tiktok:favorite:like:", value, time.Second*10)
				if err != nil {
					logc.Errorf(l.ctx, "add set to redis failed,err=%v", err)
				}
				return nil, errors.Wrapf(errorcode.FavoriteServiceDuplicateError, "user %d has liked video %d", userID, videoID)
			}

			// 没有写入缓存，先写数据库，再删除缓存
			err = l.svcCtx.VideosModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
				videoDetail, err := l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
				if err != nil {
					return err
				}

				newFavorite := favoriteModel.Favorites{
					Uid: userID,
					Vid: videoID,
				}

				_, err = l.svcCtx.FavoriteModel.Insert(l.ctx, nil, &newFavorite)
				if err != nil {
					return err
				}

				_, err = l.svcCtx.VideosModel.Update(l.ctx, nil, &videosModel.Videos{
					Id:            videoID,
					AuthorId:      videoDetail.AuthorId,
					Title:         videoDetail.Title,
					CoverUrl:      videoDetail.CoverUrl,
					PlayUrl:       videoDetail.Title,
					FavoriteCount: videoDetail.FavoriteCount + 1,
					StarCount:     videoDetail.StarCount,
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

			// 添加到redis
			_, err = l.svcCtx.BizRedis.SaddCtx(l.ctx, "tiktok:favorite:like:", value)
			if err != nil {
				return nil, errors.Wrapf(errorcode2.ServerError, "[]favorite] add set to redis failed,err=%v", err)
			}

			return nil, nil

		}

	case 2:
		// 利用redis判断存不存在
		exist, err := l.svcCtx.BizRedis.SismemberCtx(l.ctx, "tiktok:favorite:like:", value)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "redis sismember failed in FavoriteAction function,err=%v", err)
		}
		if exist {
			l.Logger.Infof("The user %d has liked the video %d and allowed to delete the like record", userID, videoID)
		} else {
			isFavorite, err := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, userID, videoID)
			if err != nil {
				if err == favoriteModel.ErrNotFound {
					logx.WithContext(l.ctx).Infof("the record is not found in the database,user_id=%d,video_id=%d", userID, videoID)
				} else {
					return nil, errors.Wrapf(errorcode2.ServerError, "find favorite record failed,err=%v", err)
				}
			}

			if !isFavorite {
				return nil, errors.Wrapf(errorcode.FavoriteServiceCancelError, "the deleted record is not exist,user_id=%d,video_id=%d", userID, videoID)
			}

			if isFavorite {
				l.svcCtx.BizRedis.SaddCtx(l.ctx, "tiktok:favorite:like:", value, time.Second*10)
			}

		}

		err = l.svcCtx.FavoriteModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			videoDetail, err := l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
			if err != nil {
				return err
			}
			id, err := l.svcCtx.FavoriteModel.FindIdByUidAndVid(l.ctx, userID, videoID)
			if err != nil {
				return err
			}
			err = l.svcCtx.FavoriteModel.Delete(l.ctx, nil, id)
			if err != nil {
				return err
			}

			_, err = l.svcCtx.VideosModel.Update(l.ctx, nil, &videosModel.Videos{
				Id:            videoID,
				AuthorId:      videoDetail.AuthorId,
				Title:         videoDetail.Title,
				CoverUrl:      videoDetail.CoverUrl,
				PlayUrl:       videoDetail.Title,
				FavoriteCount: videoDetail.FavoriteCount - 1,
				StarCount:     videoDetail.StarCount,
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

		// 添加到redis
		_, err = l.svcCtx.BizRedis.SremCtx(l.ctx, "tiktok:favorite:like:", value)
		if err != nil {
			return nil, errors.Wrapf(errorcode2.ServerError, "[]favorite] remove set element to redis failed,err=%v", err)
		}

		return nil, nil

	}
	return &favorite.FavoriteActionResponse{}, errors.Wrapf(errorcode.FavoriteInvalidActionTypeError, "the action type is invalid,action_type=%d", actionType)
}
