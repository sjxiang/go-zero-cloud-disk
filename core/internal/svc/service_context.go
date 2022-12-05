package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/config"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.DB.DataSource),
		RDB: model.InitRedis(c.Redis.Addr, c.Redis.Pass, c.Redis.DB),
	}
}
