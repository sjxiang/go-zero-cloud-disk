package upload

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
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
	
	rp := new(model.RepositoryPool)

	has, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(rp)  // 对比 repo 池的文件 hash 值，判断是否存在
	if err != nil {
		return nil, err
	}  

	resp = new(types.FileUploadPrepareResp)

	// 两种情况
	if has {
		// 1. 秒传成功
		resp.Identity = rp.Identity  // repo 池文件的 uuid
		return
	} 
	
	// 2. 获取该文件的 uploadID、key 用来进行文件的分片上传
	key, uploadId, err := util.OSSInitPart(req.Ext)
	if err != nil {
		return nil, err
	
	}
	resp.Key = key  // bucket 
	resp.UploadId = uploadId  // cos 上传凭证

	return
}
