package logic

import (
	"context"
	"errors"
	"time"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"

	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeSendLogic {
	return &VerifyCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeSendLogic) VerifyCodeSend(req *types.VerifyCodeSendReq) (resp *types.VerifyCodeSendResp, err error) {
		
	// 改邮箱未注册
	user := new(model.UserBasic)
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(user)
	if err != nil {
		return 
	}
	if cnt > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}

	// 获取验证码
	verifyCode := util.RandVerifyCode()
	expiredTime := time.Minute * time.Duration(300)  // 3 h
	
	// 存储验证码
	l.svcCtx.RDB.Set(l.ctx, req.Email, verifyCode, expiredTime)
	
	// 发送验证码
	err = util.VerifyCodeSend(req.Email, verifyCode)

	return
}
