package person_pool

import (
	"context"
	"errors"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveReq, userIdentity string) (resp *types.UserFileMoveResp, err error) {
	// todo: add your logic here and delete this line

	// 查询目录在不在，有没有这个文件夹
	parentData := new(model.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ? ", req.Parentldentity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}

	// 更新记录的 parentID
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(model.UserRepository{
		ParentId: int64(parentData.Id),  // 文件夹在 user_repo 的 id，好混乱
	})

	return
}
