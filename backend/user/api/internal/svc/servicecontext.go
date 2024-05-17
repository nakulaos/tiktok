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
	"tiktok/resource"
	"tiktok/user/api/internal/config"
	"tiktok/user/api/internal/middleware"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config                   config.Config
	Trans                    *i18n.Translator
	SetContextInfoMidlleware rest.Middleware
	AuthorityMiddleware      rest.Middleware
	UserRpc                  usesrv.UseSrv
	Cbn                      *casbin.Enforcer
	CacheConn                sqlc.CachedConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	var trans *i18n.Translator
	//if c.I18nConf.Dir != "" {
	//	trans = i18n.NewTranslatorFromFile(c.I18nConf)
	//}
	trans = i18n.NewTranslator(resource.LocaleFS)

	sqlxConn := sqlx.NewMysql(c.DB.DataSource)

	cbn := utils.NewCasbin(c.DB.DataSource, c.Casbin.Dir, c.Casbin.Table)

	return &ServiceContext{
		Config:                   c,
		Trans:                    trans,
		AuthorityMiddleware:      middleware.NewAuthorityMiddleware(cbn, sqlc.NewConn(sqlxConn, c.Cache), c.JWTPrefix, trans).Handle,
		SetContextInfoMidlleware: middleware.NewSetContextInfoMidllewareMiddleware().Handle,
		UserRpc:                  usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
		Cbn:                      utils.NewCasbin(c.DB.DataSource, c.Casbin.Dir, c.Casbin.Table),
		CacheConn:                sqlc.NewConn(sqlxConn, c.Cache),
	}
}
