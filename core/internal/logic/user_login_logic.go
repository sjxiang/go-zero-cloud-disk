package logic

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/jwt"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	
	// 1. 从数据库中查询当前用户
	user := new(model.UserBasic)

	has, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, util.MD5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或者密码错误")
	}

	// 2. 生成 token
	expiredTime, _ := strconv.Atoi(os.Getenv("TokenExpiredTIME"))
	token, err := jwt.GenerateToken(uint64(user.Id), user.Identity, user.Name, int64(expiredTime))
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{Token: token}, nil
}
