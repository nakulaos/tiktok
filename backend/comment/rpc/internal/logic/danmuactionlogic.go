package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/comment/danmumodel"
	"tiktok/comment/errorcode"
	errorcode2 "tiktok/common/errorcode"
	userModel "tiktok/user/model"

	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanMuActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDanMuActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanMuActionLogic {
	return &DanMuActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DanMuActionLogic) DanMuAction(in *comment.DanmuActionRequest) (*comment.DanmuActionResponse, error) {
	userID := in.UserId
	videoID := in.VideoId

	// 1. 判断用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.DanMuUserIdEmptyError, "user is not found , id = %d", userID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user failed, err = %v", err)
		}
	}

	// 2.检测视频是否存在
	_, err = l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.DanMuVideoIdEmptyError, "video is not found , id = %d", videoID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find video failed, err = %v", err)
		}
	}

	danmu := &danmumodel.Danmu{
		Uid:     userID,
		Vid:     videoID,
		Content: in.DanmuText,

		SendTime: float64(in.SendTime),
	}

	_, err = l.svcCtx.DanmuModel.Insert(l.ctx, nil, danmu)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "find danmu failed, err = %v", err)
	}

	return &comment.DanmuActionResponse{}, nil
}
