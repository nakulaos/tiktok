package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"tiktok/comment/commentsmodel"
	"tiktok/common/errorcode"

	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionRevertLogic {
	return &CommentActionRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionRevertLogic) CommentActionRevert(in *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	// 判断数据合法放到api中去
	l.Logger.WithContext(l.ctx).Infof("delete comment: %v", in)
	userId := in.UserId
	videoId := in.VideoId
	content := in.CommentText
	//actionType := in.ActionType
	commentId := in.CommentId

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := l.svcCtx.SqlConn.RawDB()
	if err != nil {
		return nil, errors.Wrapf(err, "db error:%v", err)
	}

	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {

		newComment := &commentsmodel.Comments{
			Uid:     userId,
			Vid:     videoId,
			Content: content,
		}

		if commentId, err = l.svcCtx.CommentsModel.DeleteWithSqlTx(l.ctx, tx, newComment); err != nil {
			return fmt.Errorf("delete comment error:%v,newComment:%v", err, newComment)
		}

		return nil
	}); err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "server error:%v", err)
	}

	return &comment.CommentActionResponse{
		CommentId: commentId,
	}, nil
}
