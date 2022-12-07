package upload

import (
	"net/http"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/logic/upload"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadPrepareReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := upload.NewFileUploadPrepareLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadPrepare(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
