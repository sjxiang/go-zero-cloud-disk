package person_pool

import (
	"context"
	"os"
	"strconv"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListReq, userIdentity string) (resp *types.UserFileListResp, err error) {
	// todo: add your logic here and delete this line
	uf := make([]*types.UserFile, 0)
	
	// 分页参数
	pageSize := req.Size
	if pageSize == 0 {
		pageSize, _ = strconv.Atoi(os.Getenv("PAGESIZE"))
	}
	page := req.Page
	if page == 0 {
		page = 1
	}

	offset := (page-1) * pageSize

	// 查询用户文件列表（直接建一张表拉到，脱裤子放屁，还连表查询）
	err = l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity,"+
				"user_repository.ext, user_repository.name, repository_pool.path, repository_pool.size" ).
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Limit(pageSize, offset).
		Find(&uf)

	if err != nil {
		return
	}

	// 计算总数
	count, err := l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(model.UserRepository))
	if err != nil {
		return
	}

	return &types.UserFileListResp{
		List: uf,
		Count: count,
	}, nil 
}
