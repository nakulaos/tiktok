package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/comment/commentsmodel"
	"tiktok/comment/errorcode"
	errorcode2 "tiktok/common/errorcode"
	userModel "tiktok/user/model"

	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrepareCommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrepareCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrepareCommentActionLogic {
	return &PrepareCommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrepareCommentActionLogic) PrepareCommentAction(in *comment.CommentActionRequest) (*comment.PrepareCommentAction, error) {
	userID := in.UserId
	videoID := in.VideoId
	actionType := in.ActionType

	// 1. 判断用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.CommentUserIdEmptyError, "user is not found , id = %d", userID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find user failed, err = %v", err)
		}
	}

	// 2.检测视频是否存在
	_, err = l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.CommentVideoIdEmptyError, "video is not found , id = %d", videoID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find video failed, err = %v", err)
		}
	}

	switch actionType {
	case 2:
		ok, err := l.svcCtx.CommentsModel.IsCommentExist(l.ctx, in.CommentId)
		if err != nil {
			if err == commentsmodel.ErrNotFound {
				return nil, errors.Wrapf(errorcode.CommentNotExistError, "comment is not found , id = %d", in.CommentId)
			}
			return nil, errors.Wrapf(errorcode2.ServerError, "server error, err = %v", err)
		}

		if !ok {
			return nil, errors.Wrapf(errorcode.CommentNotExistError, "comment is not found , id = %d", in.CommentId)
		}

	}

	return &comment.PrepareCommentAction{Ok: true}, nil
}
