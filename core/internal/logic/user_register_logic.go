package logic

import (
	"context"
	"errors"
	"log"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// 判断 code 是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取该邮箱验证码")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}

	// 判断用户名是否存在
	user := new(model.UserBasic)
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(user)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("用户名已经存在")
		return
	}

	// 数据入库
	user.Identity = util.GenUUID()
	user.Name = req.Name
	user.Password = util.MD5(req.Password)
	user.Email = req.Email

	Rows, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	
	log.Println("insert user row:", Rows)

	return
}
