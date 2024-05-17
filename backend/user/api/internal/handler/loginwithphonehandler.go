package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/user/api/internal/logic"
	"tiktok/user/api/internal/svc"
	"tiktok/user/api/internal/types"
)

func LoginWithPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginWithPhoneReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewLoginWithPhoneLogic(r.Context(), svcCtx)
		resp, err := l.LoginWithPhone(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
