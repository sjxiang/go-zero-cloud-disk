package svc

import (
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
