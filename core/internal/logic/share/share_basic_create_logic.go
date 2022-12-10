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
	
	uuid := util.GenUUID()
	
	ur := new(model.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user repo not found")
	}


	data := &model.ShareBasic {
		Identity: uuid,
		UserIdentity: userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity: ur.RepositoryIdentity,
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
