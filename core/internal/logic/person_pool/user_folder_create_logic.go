package person_pool

import (
	"context"
	"errors"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateReq, userIdentity string) (resp *types.UserFolderCreateResp, err error) {
	
	// 判断当前文件夹在该层级下是否存在
	cnt, err := l.svcCtx.Engine.
			Where("name = ? AND parent_id = ?", req.Name, req.Parentld).
			Count(new(model.UserRepository))

	if err != nil {
		return 
	}
	if cnt > 0 {
		return nil, errors.New("该文件夹已存在")
	}

	// 创建文件夹
	data := &model.UserRepository{
		Identity: util.GenUUID(),
		UserIdentity: userIdentity, 
		ParentId: req.Parentld,
		Name: req.Name,
	}

	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}

	return &types.UserFolderCreateResp{
		Identity: data.Identity,
	}, nil
}
