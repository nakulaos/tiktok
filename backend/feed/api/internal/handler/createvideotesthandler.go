package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/feed/api/internal/logic"
	"tiktok/feed/api/internal/svc"
	"tiktok/feed/api/internal/types"
)

func CreateVideoTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateVideoReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewCreateVideoTestLogic(r.Context(), svcCtx)
		err := l.CreateVideoTest(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, nil, err)
	}
}
