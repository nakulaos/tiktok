package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	KqJobConsumer        kq.KqConf
	KqUploadFileConsumer kq.KqConf

	QiNiu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		Cdn       string
		Zone      string
		Prefix    string
	}

	KqJobPush struct {
		Brokers []string
		Topic   string
	}

	KqVideoPusher struct {
		Brokers []string
		Topic   string
	}

	FeedRpc zrpc.RpcClientConf

	RecommendUrl string
}
