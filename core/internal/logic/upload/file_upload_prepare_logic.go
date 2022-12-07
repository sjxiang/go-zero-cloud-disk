package upload

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareReq) (resp *types.FileUploadPrepareResp, err error) {
	// todo: add your logic here and delete this line
	rp := new(model.RepositoryPool)

	has, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(rp)
	if err != nil {
		return 
	}  

	resp = new(types.FileUploadPrepareResp)
	if has {
		// 秒传成功
		resp.Identity = rp.Identity
	} else {
		// TODO：获取该文件的 uploadID 用来进行文件的分片上传
	}

	return
}
