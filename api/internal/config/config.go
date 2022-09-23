package config

import (
	"github.com/suyuan32/simple-admin-tools/plugins/registry/consul"
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf `yaml:",inline"`
	Auth          Auth               `json:"auth" yaml:"Auth"`
	RedisConf     redis.RedisConf    `json:"redisConf" yaml:"RedisConf"`
	CoreRpc       zrpc.RpcClientConf `json:"coreRpc" yaml:"CoreRpc"`
	Captcha       Captcha            `json:"captcha" yaml:"Captcha"`
	DB            gormsql.GORMConf   `json:"databaseConf" yaml:"DatabaseConf"`
}

type Captcha struct {
	KeyLong   int `json:"keyLong" yaml:"KeyLong"`     // captcha length
	ImgWidth  int `json:"imgWidth" yaml:"ImgWidth"`   // captcha width
	ImgHeight int `json:"imgHeight" yaml:"ImgHeight"` // captcha height
}

type Auth struct {
	AccessSecret string `json:"accessSecret" yaml:"AccessSecret"`
	AccessExpire int64  `json:"accessExpire" yaml:"AccessExpire"`
}

type ConsulConfig struct {
	Consul consul.Conf
}
