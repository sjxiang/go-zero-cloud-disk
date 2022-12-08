package person_pool

import (
	"context"
	"errors"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateReq, userIdentity string) (resp *types.UserFileNameUpdateResp, err error) {
	
	// 判断当前名称在该层级下是否存在
	cnt, err := l.svcCtx.Engine.
			Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).
			Count(new(model.UserRepository))

	if err != nil {
		return 
	}
	if cnt > 0 {
		return nil, errors.New("该文件名称已存在")
	}


	// user_repo 单方面，修改文件名称
	data := &model.UserRepository{ Name: req.Name }

	// repo 与 user_repo 数据不一致搞毛啊！
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return
	}

	return
}
