package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/favorite/api/internal/logic"
	"tiktok/favorite/api/internal/svc"
	"tiktok/favorite/api/internal/types"
)

func StarActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewStarActionLogic(r.Context(), svcCtx)
		err := l.StarAction(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, nil, err)
	}
}
