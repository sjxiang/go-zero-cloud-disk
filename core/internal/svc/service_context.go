package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/config"
	"github.com/sjxiang/go-zero-cloud-disk/core/internal/middleware"
	"github.com/sjxiang/go-zero-cloud-disk/model"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: model.Init(c.DB.DataSource),
		RDB: model.InitRedis(c.Redis.Addr, c.Redis.Pass, c.Redis.DB),
		Auth: middleware.NewAuthMiddleware().Handle,
	}
}
