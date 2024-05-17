package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Casbin struct {
		Dir   string
		Table string
	}

	DB struct {
		DataSource string
	}
	Cache cache.CacheConf

	RecommendUrl string

	JWTPrefix   string
	FavoriteRpc zrpc.RpcClientConf
}
