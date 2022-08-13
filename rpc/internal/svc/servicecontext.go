package svc

import (
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"
	"github.com/suyuan32/simple-admin-core/rpc/internal/initialize"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initialize.InitGORM(c)
	rds := initialize.InitRedis(c)
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
	}
}
