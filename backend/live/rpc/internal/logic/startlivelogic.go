package logic

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/common/errorcode"
	"tiktok/common/qiniu"

	"tiktok/live/rpc/internal/svc"
	"tiktok/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLiveLogic {
	return &StartLiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 开启直播
func (l *StartLiveLogic) StartLive(in *live.StartLiveRequest) (*live.StartLiveResponse, error) {
	idstr := strconv.Itoa(int(in.Uid))
	err := qiniu.CreatStream(idstr, l.svcCtx.Config.QiNiu.AccessKey, l.svcCtx.Config.QiNiu.SecretKey, l.svcCtx.Config.QiNiu.LiveBucket)
	if err != nil {
		return nil, errors.Wrapf(errorcode.ServerError, "create qiniu stream error:%v", err)
	}

	return &live.StartLiveResponse{
		StreamUrl: l.svcCtx.Config.QiNiu.PublishUrl + "/" + l.svcCtx.Config.QiNiu.LiveBucket + "/" + idstr,
	}, nil
}
