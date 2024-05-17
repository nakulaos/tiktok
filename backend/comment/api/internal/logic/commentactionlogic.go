package logic

import (
	"context"
	"encoding/json"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"tiktok/comment/api/internal/svc"
	"tiktok/comment/api/internal/types"
	"tiktok/comment/rpc/comment"
	"tiktok/common/errorcode"
	"tiktok/feed/rpc/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.ActionReq) (resp *types.ActionResp, err error) {
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()
	videoId := req.VideoId
	contents := req.CommentText
	actionType := req.ActionType
	commentId := req.CommentId

	l.Logger.Infof("comment action req:%+v", req)

	res, err := l.svcCtx.CommentRpc.PrepareCommentAction(l.ctx, &comment.CommentActionRequest{
		UserId:      userId,
		VideoId:     videoId,
		ActionType:  actionType,
		CommentText: contents,
		CommentId:   commentId,
	})

	l.Logger.Infof("comment action err:%v", err)

	if err != nil {
		return nil, err
	}

	var id int64

	if res.Ok {
		switch actionType {
		case 1:
			feedRpcServer, err := l.svcCtx.Config.FeedRpc.BuildTarget()
			if err != nil {
				return nil, errors.Wrapf(errorcode.ServerError, "feedrpc build target error:%v", err)
			}
			commentRpcServer, err := l.svcCtx.Config.CommentRpc.BuildTarget()
			if err != nil {
				return nil, errors.Wrapf(errorcode.ServerError, "commentRpc build target error:%v", err)
			}

			createCommentReq := &comment.CommentActionRequest{
				UserId:      userId,
				VideoId:     videoId,
				CommentText: contents,
				CommentId:   commentId,
				ActionType:  actionType,
			}

			incrCommentCountReq := &feed.IncrCommentCountReq{Id: videoId}

			gid := dtmgrpc.MustGenGid(l.svcCtx.Config.DtmServer)
			l.Logger.Infof("grpc dial %s", commentRpcServer+comment.CommentSrv_CommentAction_FullMethodName)
			saga := dtmgrpc.NewSagaGrpc(l.svcCtx.Config.DtmServer, gid).
				Add(commentRpcServer+comment.CommentSrv_CommentAction_FullMethodName,
					commentRpcServer+comment.CommentSrv_CommentActionRevert_FullMethodName,
					createCommentReq).
				Add(feedRpcServer+feed.Feed_IncrCommentCount_FullMethodName,
					feedRpcServer+feed.Feed_IncrCommentCountRevert_FullMethodName,
					incrCommentCountReq)

			err = saga.Submit()
			dtmimp.FatalIfError(err)
			if err != nil {
				return nil, errors.Wrapf(errorcode.ServerError, "submit data to  dtm-server err  : %+v \n", err)
			}

			commentidResp, err := l.svcCtx.CommentRpc.FindComment(l.ctx, &comment.CommentActionRequest{
				UserId:      userId,
				VideoId:     videoId,
				ActionType:  actionType,
				CommentText: contents,
				CommentId:   commentId,
			})
			if err != nil {
				return nil, err
			}
			id = commentidResp.Id

		case 2:
			feedRpcServer, err := l.svcCtx.Config.FeedRpc.BuildTarget()
			if err != nil {
				return nil, errors.Wrapf(errorcode.ServerError, "feedrpc build target error:%v", err)
			}
			commentRpcServer, err := l.svcCtx.Config.CommentRpc.BuildTarget()
			if err != nil {
				return nil, errors.Wrapf(errorcode.ServerError, "commentRpc build target error:%v", err)
			}

			createCommentReq := &comment.CommentActionRequest{
				UserId:      userId,
				VideoId:     videoId,
				CommentText: contents,
				CommentId:   commentId,
				ActionType:  actionType,
			}

			incrCommentCountReq := &feed.IncrCommentCountReq{Id: videoId}

			gid := dtmgrpc.MustGenGid(l.svcCtx.Config.DtmServer)
			saga := dtmgrpc.NewSagaGrpc(l.svcCtx.Config.DtmServer, gid).
				Add(commentRpcServer+comment.CommentSrv_CommentActionRevert_FullMethodName, commentRpcServer+comment.CommentSrv_CommentAction_FullMethodName,
					createCommentReq).
				Add(feedRpcServer+feed.Feed_IncrCommentCountRevert_FullMethodName, feedRpcServer+feed.Feed_IncrCommentCount_FullMethodName,
					incrCommentCountReq)

			err = saga.Submit()
			dtmimp.FatalIfError(err)
			if err != nil {
				return nil, errors.Wrapf(errorcode.ServerError, "submit data to  dtm-server err  : %+v \n", err)
			}

			id = req.CommentId
		}

	}

	return &types.ActionResp{CommentId: id}, nil
}
