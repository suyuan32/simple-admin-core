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
			logx.Errorf("DeleteInvalidTokenTask send msg error: %s\n", err.Error())
			return
		} else {
			logx.Infof("DeleteInvalidTokenTask send msg success: %s\n", res.String())
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
