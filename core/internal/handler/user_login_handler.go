package handler

import (
	"net/http"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/logic"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
