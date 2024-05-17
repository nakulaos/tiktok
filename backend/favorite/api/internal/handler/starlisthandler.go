package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/favorite/api/internal/logic"
	"tiktok/favorite/api/internal/svc"
	"tiktok/favorite/api/internal/types"
)

func StarListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewStarListLogic(r.Context(), svcCtx)
		resp, err := l.StarList(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
