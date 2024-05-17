package middleware

import (
	"encoding/json"
	"github.com/casbin/casbin/v2"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"net/http"
	"tiktok/common/base"
	"tiktok/common/errorcode"
	"tiktok/common/i18n"
	"tiktok/common/utils"
)

func AuthorityHandle(next http.HandlerFunc, cbn *casbin.Enforcer, trans *i18n.Translator, cacheConn sqlc.CachedConn, jwtPrefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj := r.URL.Path
		act := r.Method
		roldID := r.Context().Value("uid").(json.Number).String()
		cbn.LoadPolicy()
		ok, err := cbn.Enforce(roldID, obj, act)
		if err != nil {
			err = errorcode.ServerError
			err = trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}
		if !ok {
			err = errorcode.PermissionDeny
			err = trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		// 查jwt 是否过期
		var jwtResult string
		err = cacheConn.GetCacheCtx(r.Context(), jwtPrefix+utils.StripBearerPrefixFromToken(r.Header.Get("Authorization")), &jwtResult)
		if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			err = errorcode.ServerError
			err = trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		if jwtResult == "1" {
			err = errorcode.Unauthorized
			err = trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		next(w, r)
	}
}
