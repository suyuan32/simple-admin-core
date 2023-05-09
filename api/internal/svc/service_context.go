package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/captcha"

	i18n2 "github.com/suyuan32/simple-admin-core/api/internal/i18n"

	"github.com/suyuan32/simple-admin-job/job_client"

	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/rpc/core_client"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   core_client.Core
	JobRpc    job_client.Job
	Redis     *redis.Redis
	Casbin    *casbin.Enforcer
	Trans     *i18n.Translator
	Captcha   *base64Captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		CoreRpc:   core_client.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		JobRpc:    job_client.NewJob(zrpc.NewClientIfEnable(c.JobRpc)),
		Captcha:   captcha.MustNewRedisCaptcha(c.Captcha, rds),
		Redis:     rds,
		Casbin:    cbn,
		Trans:     trans,
	}
}
