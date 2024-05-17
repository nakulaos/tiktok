package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/user/api/internal/logic"
	"tiktok/user/api/internal/svc"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage()
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
