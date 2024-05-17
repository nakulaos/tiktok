package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/user/api/internal/logic"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		err := l.Register(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, nil, err)
	}
}
