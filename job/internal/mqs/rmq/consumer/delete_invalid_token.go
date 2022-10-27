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
	return &DeleteInvalidTokenTask{
		ctx:      ctx,
		svcCtx:   svcCtx,
		consumer: svcCtx.Config.ConsumerConf.NewPushConsumer(),
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
