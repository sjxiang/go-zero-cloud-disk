package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/logic"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func fileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			return 
		}

		// 判断文件是否存在
		b := make([]byte, fileHeader.Size)
		hash := fmt.Sprintf("%x", md5.Sum(b))
		
		rp := new(model.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return 
		}

		// 重复的话
		if has {  
			httpx.OkJson(w, &types.FileUploadResp{
				Identity: rp.Identity,
				Ext: rp.Ext,
				Name: rp.Name,
			})
			return
		}

		// 往 oss 中存储
		cosPath, err := util.OSSUpload(r)
		if err != nil {
			return 
		}

		// 往 logic 中传递 request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath


		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
