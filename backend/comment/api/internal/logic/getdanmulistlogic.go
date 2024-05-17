package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/comment/rpc/comment"
	"tiktok/common/errorcode"

	"tiktok/comment/api/internal/svc"
	"tiktok/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDanmuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDanmuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDanmuListLogic {
	return &GetDanmuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDanmuListLogic) GetDanmuList(req *types.DanmulistReq) (resp *types.DanmulistResp, err error) {
	res, err := l.svcCtx.CommentRpc.DanMuList(l.ctx, &comment.DanmuListRequest{VideoId: req.VideoId})
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "get danmu list error:%v", err)
	}
	resLists := make([]types.DanmuInfo, 0)

	for i := 0; i < len(res.DanmuList); i++ {

		resLists = append(resLists, types.DanmuInfo{
			Content:  res.DanmuList[i].DanmuText,
			UserId:   res.DanmuList[i].UserId,
			VideoId:  res.DanmuList[i].VideoId,
			SendTime: res.DanmuList[i].SendTime,
		})
	}
	return &types.DanmulistResp{
		DanmuList: resLists,
	}, nil
}
