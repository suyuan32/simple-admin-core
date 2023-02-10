package svc

import (
	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
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
	Redis     *redis.Redis
	Casbin    *casbin.Enforcer
	Trans     *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbin(c.DatabaseConf.Type, c.DatabaseConf.GetDSN())

	trans := &i18n.Translator{}
	trans.NewBundle(i18n.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:     rds,
		Casbin:    cbn,
		Trans:     trans,
	}
}
