package logic

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/comment/errorcode"
	errorcode2 "tiktok/common/errorcode"
	userModel "tiktok/user/model"

	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanMuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDanMuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanMuListLogic {
	return &DanMuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DanMuListLogic) DanMuList(in *comment.DanmuListRequest) (*comment.DanmuListResponse, error) {

	videoID := in.VideoId

	_, err := l.svcCtx.VideosModel.FindOne(l.ctx, videoID)
	if err != nil {
		if err == userModel.ErrNotFound {
			return nil, errors.Wrapf(errorcode.DanMuVideoIdEmptyError, "video is not found , id = %d", videoID)
		} else {
			return nil, errors.Wrapf(errorcode2.ServerError, "find video failed, err = %v", err)
		}
	}

	res, err := l.svcCtx.DanmuModel.FindDanmusByVid(l.ctx, videoID)
	if err != nil {
		return nil, errors.Wrapf(errorcode2.ServerError, "find video failed, err = %v", err)
	}

	danmuList := make([]*comment.DanMu, 0)
	for i := 0; i < len(res); i++ {
		singleinfo := &comment.DanMu{
			UserId:    int64(res[i].Uid),
			VideoId:   int64(res[i].Vid),
			DanmuText: res[i].Content,
			SendTime:  float32(res[i].SendTime),
		}
		danmuList = append(danmuList, singleinfo)
	}

	return &comment.DanmuListResponse{DanmuList: danmuList}, nil
}
