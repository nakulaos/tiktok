package handler

import (
	"net/http"
	"tiktok/comment/api/internal/logic"
	"tiktok/comment/api/internal/svc"
	"tiktok/comment/api/internal/types"
	"tiktok/common/base"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.CommentAction(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
