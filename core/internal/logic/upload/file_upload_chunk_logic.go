package upload

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkLogic {
	return &FileUploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkLogic) FileUploadChunk(req *types.FileUploadChunkReq, etag string) (resp *types.FileUploadChunkResp, err error) {
	// todo: add your logic here and delete this line

	return &types.FileUploadChunkResp{
		Etag: etag,
	}, nil
}
