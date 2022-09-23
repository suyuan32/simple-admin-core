package config

import (
	"github.com/suyuan32/simple-admin-tools/plugins/registry/consul"
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf `yaml:",inline"`
	DB                 gormsql.GORMConf `json:"DatabaseConf" yaml:"DatabaseConf"`
	RedisConf          redis.RedisConf  `json:"RedisConf" yaml:"RedisConf"`
}

type ConsulConfig struct {
	Consul consul.Conf
}
