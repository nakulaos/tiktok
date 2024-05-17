package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/feed/api/internal/logic"
	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"
)

func FindVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindVideoByIdReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewFindVideoLogic(r.Context(), svcCtx)
		resp, err := l.FindVideo(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
