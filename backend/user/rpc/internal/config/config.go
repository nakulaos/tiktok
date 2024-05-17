package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	RecommendUrl string

	Cache cache.CacheConf

	Salt string

	JWTAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Casbin struct {
		Dir   string
		Table string
	}
}
