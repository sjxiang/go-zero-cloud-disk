package logic

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq) (resp *types.FileUploadResp, err error) {
	rp := &model.RepositoryPool{
		Identity: util.GenUUID(),
		Hash: req.Hash,
		Name: req.Name,
		Ext: req.Ext,
		Size: req.Size,
		Path: req.Path,
	}

	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}

	return &types.FileUploadResp{
		Identity: rp.Identity,
		Ext: rp.Ext,
		Name: rp.Name,
	}, nil
}
