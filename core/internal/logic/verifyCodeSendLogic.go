package logic

import (
	"context"

	"github.com/sjxiang/go-zero-cloud-disk/core/internal/svc"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/types"
	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"

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
	// todo: add your logic here and delete this line
	err = util.VerifyCodeSend(req.Email, "123456")
	if err != nil {
		return nil, err
	}

	return
}
