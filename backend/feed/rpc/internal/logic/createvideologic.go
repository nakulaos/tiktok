package logic

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"tiktok/common/qiniu"
	"tiktok/feed/errorcode"
	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/internal/svc"
	videosModel "tiktok/feed/videosmodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoLogic) CreateVideo(in *feed.CreateVideoRequest) (*feed.CreateVideoResponse, error) {
	newVideo := videosModel.Videos{
		AuthorId:      int64(in.ActorId),
		Title:         in.Title,
		CoverUrl:      in.CoverUrl,
		PlayUrl:       in.Url,
		FavoriteCount: 0,
		StarCount:     0,
		CommentCount:  0,
		DelState:      0,
		Category:      int64(in.Category),
	}

	res, err := l.svcCtx.VideosModel.Insert(l.ctx, nil, &newVideo)
	if err != nil {
		return nil, errors.Wrapf(errorcode.VideoInsertError, "video:%v,err:%v", newVideo, err)
	}

	videoid, _ := res.LastInsertId()

	// 调用七牛云ai接口进行审核
	JobId := qiniu.IsSafeJobId(in.Url, strconv.Itoa(int(newVideo.Id)), l.svcCtx.Config.QiNiu.SecretKey, l.svcCtx.Config.QiNiu.AccessKey)

	// 视频信息传入mq
	jobKq := qiniu.JobBody{
		Job: JobId,
		Id:  int64(videoid),
		Url: in.Url,
		Uid: int64(newVideo.AuthorId),
	}

	// 开启子线程
	threading.GoSafe(func() {
		data, err := json.Marshal(jobKq)
		if err != nil {
			l.Logger.Errorf("[Video] marshal msg: %v error: %v", jobKq, err)
			return
		}
		err = l.svcCtx.KqJobPush.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Video] kq push data: %s error: %v", data, err)
		}
	})

	return nil, nil
}
