package config

import (
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
	LogConf   LogConf            `json:"LogConf" yaml:"LogConf"`
	RedisConf redis.RedisConf    `json:"RedisConf" yaml:"RedisConf"`
	CoreRpc   zrpc.RpcClientConf `json:"CoreRpc" yaml:"CoreRpc"`
	Captcha   Captcha            `json:"Captcha" yaml:"Captcha"`
	DB        gormsql.GORMConf   `json:"DatabaseConf" yaml:"DatabaseConf"`
}

type LogConf struct {
	ServiceName         string `json:",optional"`                                    // service name
	Mode                string `json:",default=console,options=console|file|volume"` // mode: console-output to console，file-output to file，，volume-output to the docker volume
	Path                string `json:",default=logs"`                                // store path
	Level               string `json:",default=info,options=info|error|severe"`      // the level to be shown
	Compress            bool   `json:",optional"`                                    // gzip compress
	KeepDays            int    `json:",optional"`                                    // the period to be stored
	StackCooldownMillis int    `json:",default=100"`                                 // the period between two writing (ms)
}

type NodeConf struct {
	redis.RedisConf
	Weight int `json:",default=100"` // weight , default is 100
}

type Captcha struct {
	KeyLong   int `json:"KeyLong" yaml:"KeyLong"`     // captcha length
	ImgWidth  int `json:"ImgWidth" yaml:"ImgWidth"`   // captcha width
	ImgHeight int `json:"ImgHeight" yaml:"ImgHeight"` // captcha height
}
