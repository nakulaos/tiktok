package handler

import (
	"net/http"
	"tiktok/comment/api/internal/logic"
	"tiktok/comment/api/internal/svc"
	"tiktok/comment/api/internal/types"
	"tiktok/common/base"
)

func DanmuActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DanmuActionReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewDanmuActionLogic(r.Context(), svcCtx)
		resp, err := l.DanmuAction(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
