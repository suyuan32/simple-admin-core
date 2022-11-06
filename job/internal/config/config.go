package config

import (
	"github.com/zeromicro/go-zero/core/service"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/plugins/mq/rocketmq"

	"github.com/suyuan32/simple-admin-core/pkg/config"
)

type Config struct {
	service.ServiceConf

	DatabaseConf config.DatabaseConf
	Redis        redis2.RedisConf
	ConsumerConf rocketmq.ConsumerConf
	ProducerConf rocketmq.ProducerConf
}
