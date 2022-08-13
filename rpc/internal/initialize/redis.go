package initialize

import (
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func InitRedis(c config.Config) *redis.Redis {
	r := redis.New(c.RedisConf.Host, func(r *redis.Redis) {
		r.Type = c.RedisConf.Type
		r.Pass = c.RedisConf.Pass
	})
	return r
}
