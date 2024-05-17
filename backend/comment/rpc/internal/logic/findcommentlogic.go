package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/comment/commentsmodel"
	"tiktok/comment/errorcode"
	errorcode2 "tiktok/common/errorcode"

	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentLogic {
	return &FindCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindCommentLogic) FindComment(in *comment.CommentActionRequest) (*comment.FindCommentResp, error) {
	id, err := l.svcCtx.CommentsModel.FindComment(l.ctx, in.UserId, in.VideoId, in.CommentText)
	if err != nil {
		if err == commentsmodel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.CommentNotExistError, "comment is not found , uid =%d vid =%d content:%s", in.UserId, in.VideoId, in.CommentText)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find comment failed, err = %v", err)
		}
	}
	return &comment.FindCommentResp{Id: id}, nil
}
