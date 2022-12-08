package upload

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteReq) (resp *types.FileUploadChunkCompleteResp, err error) {
	// todo: add your logic here and delete this line

	co := make([]cos.Object, 0)

	for _, v := range req.CosObjects {
		co = append(co, cos.Object{
			ETag: v.Etag,
			PartNumber: v.PartNumber,
		})
	}

	err = util.OSSPartUploadComplete(req.Key, req.UploadId, co)
	if err != nil {
		return 
	}

	return
}
