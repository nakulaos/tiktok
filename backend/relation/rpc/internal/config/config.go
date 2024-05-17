package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Salt       string
	DB         struct {
		DataSource string
	}
	UserRpc zrpc.RpcClientConf
	Cache   cache.CacheConf
}
