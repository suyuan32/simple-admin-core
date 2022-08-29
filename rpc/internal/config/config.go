package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB        gormsql.GORMConf `json:"DatabaseConf" yaml:"DatabaseConf"`
	LogConf   logx.LogConf     `json:"LogConf" yaml:"LogConf"`
	RedisConf redis.RedisConf  `json:"RedisConf" yaml:"RedisConf"`
}
