package svc

import (
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"
	"gorm.io/gorm"
	"log"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := c.DB.NewGORM()
	if err != nil {
		log.Fatal(err.Error())
	}
	rds := c.RedisConf.NewRedis()
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
	}
}
