package upload

import (
	"net/http"
	"fmt"
	"crypto/md5"
	"path"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/logic/upload"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// === start
		// formdata 读数据
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			return 
		}

		// 判断文件在 repo 池中，是否存在
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

		// 否则，往 oss 中存储
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


		// === over

		l := upload.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
