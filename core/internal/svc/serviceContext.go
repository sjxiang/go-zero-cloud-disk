package svc

import (
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/config"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.DB.DataSource),
	}
}
