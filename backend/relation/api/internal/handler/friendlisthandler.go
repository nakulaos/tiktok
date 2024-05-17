package handler

import (
	"net/http"
	"tiktok/common/base"
	"tiktok/relation/api/internal/logic"
	"tiktok/relation/api/internal/svc"
	"tiktok/relation/api/internal/types"
)

func FriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendListReq
		if err := base.Parse(r, &req); err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r, w, nil, err)
			return
		}

		l := logic.NewFriendListLogic(r.Context(), svcCtx)
		resp, err := l.FriendList(&req)
		err = svcCtx.Trans.TransError(r.Context(), err)
		base.HttpResult(r, w, resp, err)
	}
}
