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
	UserRpc zrpc.RpcClientConf
	Cache   cache.CacheConf
	QiNiu   struct {
		AccessKey    string
		SecretKey    string
		Bucket       string
		Cdn          string
		Zone         string
		Prefix       string
		LiveBucket   string
		PublishUrl   string
		LiveUrl      string
		LiveCoverUrl string
	}
}
