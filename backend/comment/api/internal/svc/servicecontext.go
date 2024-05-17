package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tiktok/comment/api/internal/config"
	"tiktok/comment/api/internal/middleware"
	"tiktok/comment/rpc/comment"
	"tiktok/comment/rpc/commentsrv"
	"tiktok/common/i18n"
	"tiktok/common/interceptors"
	"tiktok/common/utils"
	"tiktok/feed/rpc/feed"
	"tiktok/feed/rpc/feedclient"
	"tiktok/resource"
	"tiktok/user/rpc/user"
	"tiktok/user/rpc/usesrv"
)

type ServiceContext struct {
	Config                   config.Config
	Trans                    *i18n.Translator
	SetContextInfoMidlleware rest.Middleware
	AuthorityMiddleware      rest.Middleware
	Cbn                      *casbin.Enforcer
	CacheConn                sqlc.CachedConn
	FeedRpc                  feed.FeedClient
	CommentRpc               comment.CommentSrvClient
	UserRpc                  user.UseSrvClient
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
		FeedRpc:                  feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
		CommentRpc:               commentsrv.NewCommentSrv(zrpc.MustNewClient(c.CommentRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
		UserRpc:                  usesrv.NewUseSrv(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))),
	}
}
