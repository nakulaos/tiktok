package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	DB         struct {
		DataSource string
	}
	Cache    cache.CacheConf
	BizRedis redis.RedisConf
	UserRpc  zrpc.RpcClientConf
}
