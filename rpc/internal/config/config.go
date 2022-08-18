package config

import (
	"github.com/zeromicro/go-zero/core/stores/gormsql"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB        gormsql.GORMConf `json:"DatabaseConf" yaml:"DatabaseConf"`
	LogConf   LogConf          `json:"LogConf" yaml:"LogConf"`
	RedisConf redis.RedisConf  `json:"RedisConf" yaml:"RedisConf"`
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
