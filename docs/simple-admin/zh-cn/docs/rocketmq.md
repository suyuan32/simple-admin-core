# 消息队列

> 我们使用 rocketmq 实现消息队列

> Producer

将 producer 的任务添加到 job/internal/mqs/producer 中， 参考

```go
package producer

import (
	"context"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/go-co-op/gocron"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/job/internal/svc"
)

type DeleteInvalidTokenTask struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	producer rocketmq.Producer
	cron     *gocron.Scheduler
}

func NewDeleteInvalidTokenTask(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInvalidTokenTask {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(svcCtx.Config.ProducerConf.NsResolver)),
		producer.WithRetry(svcCtx.Config.ProducerConf.Retry))
	logx.Must(err)
	return &DeleteInvalidTokenTask{
		ctx:      ctx,
		svcCtx:   svcCtx,
		producer: p,
		cron:     gocron.NewScheduler(time.UTC),
	}
}

func (l *DeleteInvalidTokenTask) Start() {
	logx.Info("DeleteInvalidTokenTask producer start")
	err := l.producer.Start()
	logx.Must(err)

	// delete invalid token every 1 minute
	_, err = l.cron.Every(1).Minute().Do(func() {
		msg := &primitive.Message{
			Topic: "delete-invalid-token",
			Body:  []byte("all"),
		}
		msg.WithKeys([]string{"DeleteInvalidTokenTask"})

		res, err := l.producer.SendSync(context.Background(), msg)

		if err != nil {
			logx.Errorf("DeleteInvalidTokenTask send message error: %s\n", err.Error())
			return
		} else {
			logx.Infof("DeleteInvalidTokenTask send message success: %s\n", res.String())
		}
	})

	if err != nil {
		logx.Error("producer error: %s\n", err.Error())
		return
	}

	l.cron.StartAsync()
}

func (l *DeleteInvalidTokenTask) Stop() {
	err := l.producer.Shutdown()
	if err != nil {
		logx.Errorw("DeleteInvalidTokenTask producer cannot shut down")
		return
	}
	l.cron.Stop()
}

```

> Consumer 

将 Consumer 的任务添加到 job/internal/mqs/consumer 中， 参考

```go
package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-co-op/gocron"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/job/internal/svc"
)

type DeleteInvalidTokenTask struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	consumer rocketmq.PushConsumer
	cron     *gocron.Scheduler
}

func NewDeleteInvalidTokenTask(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInvalidTokenTask {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(svcCtx.Config.ConsumerConf.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(svcCtx.Config.ConsumerConf.NsResolver)))
	logx.Must(err)

	return &DeleteInvalidTokenTask{
		ctx:      ctx,
		svcCtx:   svcCtx,
		consumer: c,
		cron:     gocron.NewScheduler(time.UTC),
	}
}

func (l *DeleteInvalidTokenTask) Start() {
	logx.Info("DeleteInvalidTokenTask consumer start")
	err := l.consumer.Start()
	logx.Must(err)

	// delete invalid token every 1 minute
	l.consumer.Subscribe("delete-invalid-token", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}

		return consumer.ConsumeSuccess, nil
	})
}

func (l *DeleteInvalidTokenTask) Stop() {
	err := l.consumer.Shutdown()
	if err != nil {
		logx.Errorw("DeleteInvalidTokenTask consumer cannot shut down")
		return
	}
	l.cron.Stop()
}

```

>注意： 需要实现 Start 和 Stop 方法

> 添加监听

修改 job/internal/listen/rmq.go

```go
package listen

import (
	"context"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/suyuan32/simple-admin-core/job/internal/config"
	"github.com/suyuan32/simple-admin-core/job/internal/mqs/rmq/consumer"
	"github.com/suyuan32/simple-admin-core/job/internal/mqs/rmq/producer"
	"github.com/suyuan32/simple-admin-core/job/internal/svc"
)

// Rmq RocketMQ service
func Rmq(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		consumer.NewDeleteInvalidTokenTask(ctx, svcCtx),
		producer.NewDeleteInvalidTokenTask(ctx, svcCtx),
	}
}

```

> 启动服务

在 job 文件夹执行

```shell
go run core.go -f etc/core.yaml
```

> 集成 producer 进 rpc 或 api

需要在 service_context.go 初始化全局变量， 还需要在 config 添加配置

```go
package svc

import (
	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
	Producer rocketmq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize database connection
	db, err := c.DatabaseConf.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successfully")
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successfully")

	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(svcCtx.Config.ProducerConf.NsResolver)),
		producer.WithRetry(svcCtx.Config.ProducerConf.Retry))
	logx.Must(err)
	
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
		Producer: p,
	}
}

```