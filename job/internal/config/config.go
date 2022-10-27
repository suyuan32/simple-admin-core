package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/gormsql"
	"github.com/zeromicro/go-zero/plugins/mq/rocketmq"
)

type Config struct {
	service.ServiceConf

	DatabaseConf gormsql.GORMConf
	ConsumerConf rocketmq.ConsumerConf
	ProducerConf rocketmq.ProducerConf
}
