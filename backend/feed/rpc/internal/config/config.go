package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	DB         struct {
		DataSource string
	}
	Cache cache.CacheConf

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

	UserRpc      zrpc.RpcClientConf
	RecommendUrl string

	Es struct {
		Addresses []string
		//Username   string
		//Password   string
		//MaxRetries int
	}
}
