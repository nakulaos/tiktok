package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"tiktok/comment/commentsmodel"
	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"
	"tiktok/common/errorcode"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	// 判断数据合法放到api中去
	l.Logger.WithContext(l.ctx).Infof("create comment: %v", in)
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

		if commentId, err = l.svcCtx.CommentsModel.InsertWithSqlTx(l.ctx, tx, newComment); err != nil {
			return fmt.Errorf("insert comment error:%v,newComment:%v", err, newComment)
		}

		return nil
	}); err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "server error:%v", err)
	}

	// 判断数据过程放到api中去

	return &comment.CommentActionResponse{CommentId: commentId}, nil
}
