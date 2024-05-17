package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/common/errorcode"

	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoDurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoDurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoDurationLogic {
	return &VideoDurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoDurationLogic) VideoDuration(in *feed.VideoDurationReq) (*feed.VideoDurationResp, error) {
	video, err := l.svcCtx.VideosModel.FindOne(l.ctx, int64(in.VideoId))
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "find video error:%v", err)
	}

	video.Duration.String = in.Duration
	video.Duration.Valid = true

	_, err = l.svcCtx.VideosModel.Update(l.ctx, nil, video)

	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "update video error:%v", err)
	}

	return &feed.VideoDurationResp{}, nil
}
