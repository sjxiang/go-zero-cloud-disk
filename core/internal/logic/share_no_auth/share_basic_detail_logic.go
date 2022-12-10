package share_no_auth

import (
	"context"
	"errors"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailReq) (resp *types.ShareBasicDetailResp, err error) {
	
		// 对分享记录的点击次数进行 +1 
		_, err = l.svcCtx.Engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
		if err != nil {
			return nil, errors.New("更新点击次数失败，err:" + err.Error())
		}
		
	
		
		// 获取资源的详细信息（其实最好是用户自己 repo，而不是中心 repo 数据）
		resp = new(types.ShareBasicDetailResp)
		l.svcCtx.Engine.Logger().ShowSQL(true)
	
		_, err = l.svcCtx.Engine.Table("share_basic").
			Select("share_basic.repository_identity,"+
				   "user_repository.name,"+
				   "repository_pool.ext,"+
				   "repository_pool.size,"+
				   "repository_pool.path").
			Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
			Join("LEFT", "user_repository", "share_basic.user_repository_identity = user_repository.identity").
			Where("share_basic.identity = ?", req.Identity).
			Get(resp)
			
		if err != nil {
			return nil, errors.New("连接查询失败，err：" + err.Error())
		}
	
		return resp, nil
}
