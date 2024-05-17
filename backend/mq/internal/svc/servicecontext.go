package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/common/interceptors"
	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/feedclient"
	"tiktok/mq/internal/config"
)

type ServiceContext struct {
	Config               config.Config
	KqJobConsumer        queue.MessageQueue
	KqUploadFileConsumer queue.MessageQueue
	KqJobPush            *kq.Pusher
	FeedRpc              feed.FeedClient
	KqVideoPusher        *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		KqJobPush:     kq.NewPusher(c.KqJobPush.Brokers, c.KqJobPush.Topic, kq.WithAllowAutoTopicCreation()),
		FeedRpc:       feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
		KqVideoPusher: kq.NewPusher(c.KqVideoPusher.Brokers, c.KqVideoPusher.Topic, kq.WithAllowAutoTopicCreation()),
	}

}
