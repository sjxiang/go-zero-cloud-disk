package share

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateReq, userIdentity string) (resp *types.ShareBasicCreateResp, err error) {
	// todo: add your logic here and delete this line
	
	data := &model.ShareBasic {
		Identity: util.GenUUID(),
		UserIdentity: userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ExpiredTime: req.ExpiredTime,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}

	return &types.ShareBasicCreateResp{
		Identity: data.Identity,
	}, nil
}
