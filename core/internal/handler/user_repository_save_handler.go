package handler

import (
	"net/http"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/logic"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userRepositorySaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepositorySaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRepositorySaveLogic(r.Context(), svcCtx)
		resp, err := l.UserRepositorySave(&req, r.Header.Get("UserIdentity"))  // 这边是分离了 header 和 body
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
