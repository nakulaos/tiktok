package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/user/api/internal/logic"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, nil, err)
	}
}
