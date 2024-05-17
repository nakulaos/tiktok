package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	ii18n "tiktok/common/i18n"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	I18nConf ii18n.Conf

	Casbin struct {
		Dir   string
		Table string
	}

	DB struct {
		DataSource string
	}
	Cache cache.CacheConf

	RecommendUrl string

	JWTPrefix string

	RelationRpc zrpc.RpcClientConf
}
