package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    redis.RedisConf
	CoreRpc      zrpc.RpcClientConf
	JobRpc       zrpc.RpcClientConf
	Captcha      captcha.Conf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	I18nConf     i18n.Conf
	ProjectConf  ProjectConf
	CROSConf     config.CROSConf
}

type ProjectConf struct {
	DefaultRole       uint64 `json:",default=1"`
	DefaultDepartment uint64 `json:",default=1"`
	DefaultPosition   uint64 `json:",default=1"`
}
