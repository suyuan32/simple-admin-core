package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    config.RedisConf
	CoreRpc      zrpc.RpcClientConf
	JobRpc       zrpc.RpcClientConf
	McmsRpc      zrpc.RpcClientConf
	Captcha      captcha.Conf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	I18nConf     i18n.Conf
	ProjectConf  ProjectConf
	CROSConf     config.CROSConf
}

type ProjectConf struct {
	DefaultRoleId           uint64 `json:",default=1"`
	DefaultRegisterHomePath string `json:",default=/dashboard"`
	DefaultDepartmentId     uint64 `json:",default=1"`
	DefaultPositionId       uint64 `json:",default=1"`
	EmailCaptchaExpiredTime int    `json:",default=600"`
	SmsTemplateId           string `json:",optional"`
	SmsAppId                string `json:",optional"`
	SmsSignName             string `json:",optional"`
	SmsParamsType           string `json:",default=json,options=[json,array]"`
	RegisterVerify          string `json:",default=captcha,options=[disable,captcha,email,sms,sms_or_email]"`
	LoginVerify             string `json:",default=captcha,options=[captcha,email,sms,sms_or_email,all]"`
	ResetVerify             string `json:",default=email,options=[email,sms,sms_or_email]"`
	AllowInit               bool   `json:",default=true"`
	RefreshTokenPeriod      int    `json:",optional,default=24"` // refresh token valid period, unit: hour | 刷新 token 的有效期，单位：小时
	AccessTokenPeriod       int    `json:",optional,default=1"`  // access token valid period, unit: hour | 短期 token 的有效期，单位：小时
}
