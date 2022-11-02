package svc

import (
	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	Redis     *redis.Redis
	Casbin    *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successfully")

	// initialize database connection
	db, err := c.DatabaseConf.NewGORM()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successful")

	// initialize casbin
	cbn := utils.NewCasbin(db)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:     rds,
		Casbin:    cbn,
	}
}
