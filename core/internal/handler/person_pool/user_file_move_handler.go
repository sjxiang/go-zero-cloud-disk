package person_pool

import (
	"net/http"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/logic/person_pool"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileMoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileMoveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := person_pool.NewUserFileMoveLogic(r.Context(), svcCtx)
		resp, err := l.UserFileMove(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
