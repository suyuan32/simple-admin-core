package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-core/pkg/config"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	CasbinConf   config.CasbinConf
	RedisConf    redis.RedisConf
	GlobalEnv    config.GLobalEnv
}
