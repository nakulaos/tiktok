package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/relation/api/internal/logic"
	"tiktok/relation/api/internal/svc"
	"tiktok/relation/api/internal/types"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewFavoriteActionLogic(r.Context(), svcCtx)
		err := l.FavoriteAction(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, nil, err)
	}
}
