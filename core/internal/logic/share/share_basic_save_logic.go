package share

import (
	"context"
	"errors"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveReq, userIdentity string) (resp *types.ShareBasicSaveResp, err error) {
	
	// 获取资源的详情
	rp := new(model.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("资源不存在")
	}

	// user_repo 资源保存
	ur := &model.UserRepository{
		Identity: util.GenUUID(),
		UserIdentity: userIdentity,
		ParentId: req.Parentld,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext: rp.Ext,
		Name: rp.Name,
	}

	_, err = l.svcCtx.Engine.Insert(ur)

	return &types.ShareBasicSaveResp{
		Identity: ur.Identity,
	}, nil
}
