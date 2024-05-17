package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/common/i18n"
	"tiktok/common/interceptors"
	"tiktok/common/utils"
	"tiktok/favorite/api/internal/config"
	"tiktok/favorite/api/internal/middleware"
	"tiktok/favorite/rpc/favoriteclient"
	"tiktok/resource"
)

type ServiceContext struct {
	Config                   config.Config
	Trans                    *i18n.Translator
	SetContextInfoMidlleware rest.Middleware
	AuthorityMiddleware      rest.Middleware
	Cbn                      *casbin.Enforcer
	CacheConn                sqlc.CachedConn
	FavoriteRpc              favoriteclient.Favorite
}

func NewServiceContext(c config.Config) *ServiceContext {
	var trans *i18n.Translator

	trans = i18n.NewTranslator(resource.LocaleFS)

	sqlxConn := sqlx.NewMysql(c.DB.DataSource)

	cbn := utils.NewCasbin(c.DB.DataSource, c.Casbin.Dir, c.Casbin.Table)

	return &ServiceContext{
		Config:                   c,
		Trans:                    trans,
		AuthorityMiddleware:      middleware.NewAuthorityMiddleware(cbn, sqlc.NewConn(sqlxConn, c.Cache), c.JWTPrefix, trans).Handle,
		SetContextInfoMidlleware: middleware.NewSetContextInfoMidllewareMiddleware().Handle,
		Cbn:                      utils.NewCasbin(c.DB.DataSource, c.Casbin.Dir, c.Casbin.Table),
		CacheConn:                sqlc.NewConn(sqlxConn, c.Cache),
		FavoriteRpc:              favoriteclient.NewFavorite(zrpc.MustNewClient(c.FavoriteRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
	}
}
