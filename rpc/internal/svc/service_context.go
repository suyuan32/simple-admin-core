package svc

import (
	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize database connection
	db, err := c.DatabaseConf.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successfully")
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successfully")
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
	}
}
