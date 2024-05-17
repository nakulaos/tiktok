package logic

import (
	"context"
	"tiktok/comment/rpc/comment"

	"tiktok/comment/api/internal/svc"
	"tiktok/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuActionLogic {
	return &DanmuActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuActionLogic) DanmuAction(req *types.DanmuActionReq) (resp *types.DanmuActionResp, err error) {
	_, err = l.svcCtx.CommentRpc.DanMuAction(l.ctx, &comment.DanmuActionRequest{
		UserId:    req.Author,
		VideoId:   req.VideoId,
		DanmuText: req.DanmuText,
		SendTime:  req.SendTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.DanmuActionResp{}, nil
}
