package person_pool

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveReq, userIdentity string) (resp *types.UserRepositorySaveResp, err error) {
	
	// user_repo 保存一份
	ur := &model.UserRepository{
		Identity: util.GenUUID(),
		UserIdentity: userIdentity,  // jwt -> header -> user_repo
		ParentId: req.ParentIld,
		RepositoryIdentity: req.RepositoryIdentity,  // repo && user_repo
		Ext: req.Ext,
		Name: req.Name,
	}

	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return 
	}

	return &types.UserRepositorySaveResp{
		Identity: ur.Identity,
	}, nil
}
