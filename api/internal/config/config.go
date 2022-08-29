package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	LogConf   logx.LogConf       `json:"LogConf" yaml:"LogConf"`
	RedisConf redis.RedisConf    `json:"RedisConf" yaml:"RedisConf"`
	CoreRpc   zrpc.RpcClientConf `json:"CoreRpc" yaml:"CoreRpc"`
	Captcha   Captcha            `json:"Captcha" yaml:"Captcha"`
	DB        gormsql.GORMConf   `json:"DatabaseConf" yaml:"DatabaseConf"`
}

type Captcha struct {
	KeyLong   int `json:"KeyLong" yaml:"KeyLong"`     // captcha length
	ImgWidth  int `json:"ImgWidth" yaml:"ImgWidth"`   // captcha width
	ImgHeight int `json:"ImgHeight" yaml:"ImgHeight"` // captcha height
}
