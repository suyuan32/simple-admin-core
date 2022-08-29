package svc

import (
	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/core"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/logx/zapx"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   core.Core
	Redis     *redis.Redis
	Casbin    *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize logx
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	logx.MustSetup(c.LogConf)
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successful")
	// initialize database connection
	db, err := c.DB.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successful")
	cbn := utils.NewCasbin(db)
	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		CoreRpc:   core.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:     rds,
		Casbin:    cbn,
	}
}
