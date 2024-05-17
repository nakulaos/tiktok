package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"net/http"
	"tiktok/common/i18n"
	"tiktok/common/middleware"
)

type AuthorityMiddleware struct {
	cbn       *casbin.Enforcer
	cacheConn sqlc.CachedConn
	jwtPrefix string
	trans     *i18n.Translator
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, cacheConn sqlc.CachedConn, jwtPrefix string, trans *i18n.Translator) *AuthorityMiddleware {
	return &AuthorityMiddleware{cbn: cbn, cacheConn: cacheConn, jwtPrefix: jwtPrefix, trans: trans}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return middleware.AuthorityHandle(next, m.cbn, m.trans, m.cacheConn, m.jwtPrefix)
}
