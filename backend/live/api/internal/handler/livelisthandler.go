package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/live/api/internal/logic"
	"tiktok/live/api/internal/svc"
)

func LiveListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLiveListLogic(r.Context(), svcCtx)
		resp, err := l.LiveList()
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
