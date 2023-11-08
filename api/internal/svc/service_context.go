package svc

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	i18n2 "github.com/suyuan32/simple-admin-core/api/internal/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/suyuan32/simple-admin-message-center/mcmsclient"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"

	"github.com/suyuan32/simple-admin-job/jobclient"

	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Authority   rest.Middleware
	CoreRpc     coreclient.Core
	JobRpc      jobclient.Job
	McmsRpc     mcmsclient.Mcms
	Redis       *redis.Redis
	Casbin      *casbin.Enforcer
	Trans       *i18n.Translator
	Captcha     *base64Captcha.Captcha
	BanRoleData map[string]bool // ban role means the role status is not normal
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(),
		c.RedisConf)

	var trans *i18n.Translator
	if c.I18nConf.Dir != "" {
		trans = i18n.NewTranslatorFromFile(c.I18nConf)
	} else {
		trans = i18n.NewTranslator(i18n2.LocaleFS)
	}

	svc := &ServiceContext{
		Config:  c,
		CoreRpc: coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		JobRpc:  jobclient.NewJob(zrpc.NewClientIfEnable(c.JobRpc)),
		McmsRpc: mcmsclient.NewMcms(zrpc.NewClientIfEnable(c.McmsRpc)),
		Captcha: captcha.MustNewRedisCaptcha(c.Captcha, rds),
		Redis:   rds,
		Casbin:  cbn,
		Trans:   trans,
	}

	err := svc.LoadBanRoleData()
	logx.Must(err)

	svc.Authority = middleware.NewAuthorityMiddleware(cbn, rds, trans, svc.BanRoleData).Handle

	return svc
}

func (l *ServiceContext) LoadBanRoleData() error {
	if !l.Config.CoreRpc.Enabled {
		return errorx.NewCodeInternalError(i18n.ServiceUnavailable)
	}

	if l.BanRoleData == nil {
		l.BanRoleData = make(map[string]bool)
	}

	roleData, err := l.CoreRpc.GetRoleList(context.Background(), &core.RoleListReq{
		Page:     1,
		PageSize: 1000,
	})

	if err != nil {
		if strings.Contains(err.Error(), i18n.DatabaseError) {
			return nil
		}
		logx.Error("failed to load role data, please check if initialize the database")
		return errorx.NewCodeInternalError("failed to load role data")
	}

	var state bool
	for _, v := range roleData.Data {
		if uint8(*v.Status) == common.StatusNormal {
			state = false
		} else {
			state = true
		}

		l.BanRoleData[*v.Code] = state
	}

	return nil
}
