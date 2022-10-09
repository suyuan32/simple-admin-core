package config

import (
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB        gormsql.GORMConf `json:"DatabaseConf" yaml:"DatabaseConf"`
	RedisConf redis.RedisConf  `json:"RedisConf" yaml:"RedisConf"`
}
