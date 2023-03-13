package svc

import (
	"github.com/suyuan32/simple-admin-common/i18n"

	i18n2 "github.com/suyuan32/simple-admin-core/api/internal/i18n"

	"github.com/suyuan32/simple-admin-job/jobclient"

	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	JobRpc    jobclient.Job
	Redis     *redis.Redis
	Casbin    *casbin.Enforcer
	Trans     *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		JobRpc:    jobclient.NewJob(zrpc.NewClientIfEnable(c.JobRpc)),
		Redis:     rds,
		Casbin:    cbn,
		Trans:     trans,
	}
}
