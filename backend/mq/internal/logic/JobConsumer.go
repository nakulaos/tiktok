package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"tiktok/common/qiniu"
	"tiktok/feed/rpc/feed"
	"tiktok/mq/internal/svc"
	"time"
)

type JobConsumerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobConsumerLogic {
	return &JobConsumerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobConsumerLogic) Consume(key, val string) error {
	l.Logger.Infof("Job key: %s, val: %s", key, val)
	var job qiniu.JobBody
	err := json.Unmarshal([]byte(val), &job)
	if err != nil {
		l.Logger.Errorf("json unmarshal error: %v", err)
		return err
	}

	status, vid, suggestion, err := qiniu.GetJobBack(job.Job, l.svcCtx.Config.QiNiu.SecretKey, l.svcCtx.Config.QiNiu.AccessKey)
	if err != nil {
		l.Logger.Errorf("job back ,job :%v,error: %v", job, err)
		return err
	}

	if status != "FINISHED" {
		// 未完成或者失败，重新放进队列
		Jobkq := qiniu.JobBody{
			Job: job.Job,
			Id:  job.Id,
			Url: job.Url,
			Uid: job.Uid,
		}
		time.Sleep(5 * time.Second)
		data, err := json.Marshal(Jobkq)
		if err != nil {
			l.Logger.Errorf("json marshal error: %v", err)
			return err
		}
		err = l.svcCtx.KqJobPush.Push(string(data))
		if err != nil {
			return err
		}
	} else {
		// 审核未通过
		if suggestion != "pass" {
			num, _ := strconv.ParseInt(vid, 10, 32)
			ctx, _ := context.WithTimeout(l.ctx, time.Second*15)
			video, err := l.svcCtx.FeedRpc.DeleteVideo(ctx, &feed.DeleteVideoReq{Vid: int32(num)})
			if err != nil {
				l.Logger.Errorf("video delete error: %v", err)
				return err
			}
			l.Logger.Infof("video delete success: %v", video)
			return nil
		} else {
			messagekq := qiniu.UploadFile{
				Id:  job.Id,
				Url: job.Url,
				Uid: job.Uid,
			}
			// 发送kafka消息，异步
			threading.GoSafe(func() {
				data, err := json.Marshal(messagekq)
				if err != nil {
					return
				}
				err = l.svcCtx.KqVideoPusher.Push(string(data))
			})
			l.Logger.Infof("video upload success: %v", messagekq)
			return nil
		}
	}
	return nil
}
