package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/zeromicro/go-zero/core/service"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/plugins/mq/rocketmq"
)

type Config struct {
	service.ServiceConf

	DatabaseConf config.DatabaseConf
	RedisConf    redis2.RedisConf
	ConsumerConf rocketmq.ConsumerConf
	ProducerConf rocketmq.ProducerConf
}
