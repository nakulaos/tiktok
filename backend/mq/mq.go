package main

import (
	"flag"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/net/context"
	"tiktok/mq/internal/config"
	"tiktok/mq/internal/logic"
	"tiktok/mq/internal/svc"
)

var configFile = flag.String("f", "etc/mq_dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svcCtx := svc.NewServiceContext(c)
	jobConsumerLogic := logic.NewJobConsumerLogic(context.Background(), svcCtx)
	uploadFileLogic := logic.NewUploadFileLogic(context.Background(), svcCtx)

	svcCtx.KqJobConsumer = kq.MustNewQueue(c.KqJobConsumer, jobConsumerLogic)
	svcCtx.KqUploadFileConsumer = kq.MustNewQueue(c.KqUploadFileConsumer, uploadFileLogic)

	//svcCtx.KqJobConsumer.Start()
	//go func() {
	//	svcCtx.KqUploadFileConsumer.Start()
	//}()

	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(svcCtx.KqJobConsumer)
	serviceGroup.Add(svcCtx.KqUploadFileConsumer)

	serviceGroup.Start()

	defer serviceGroup.Stop()

	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//server.Start()
	//
	//defer svcCtx.KqJobConsumer.Stop()
	//defer svcCtx.KqUploadFileConsumer.Stop()
}
