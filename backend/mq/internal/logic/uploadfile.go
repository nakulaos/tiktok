package logic

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"tiktok/common/gorse"
	"tiktok/common/qiniu"
	"tiktok/common/utils"
	"tiktok/mq/internal/svc"
	"time"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileLogic) Consume(key, val string) error {
	logx.Infof("upload file key :%s , val :%s", key, val)
	var videoInfo qiniu.UploadFile
	err := json.Unmarshal([]byte(val), &videoInfo)
	if err != nil {
		fmt.Println("unmarshal json failed:%s\n", err.Error())
		return err
	}

	// TODO: ai接口实现

	go func() {
		var labels []string
		postbaseurl := l.svcCtx.Config.RecommendUrl + "/api/item"
		requestBody := gorse.VideosGoresBody{
			ItemId:    fmt.Sprintf("%d", videoInfo.Id),
			Timestamp: time.Now(),
			Labels:    labels,
		}

		// 将请求体编码为JSON字节
		fmt.Printf("url:%s,requestBody:%v\n", postbaseurl, requestBody)
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("um Marshal:", err)
			return
		}
		post, err := utils.Post(postbaseurl, jsonData)
		// 打印响应内容
		fmt.Println(string(post))

	}()

	return nil

}
