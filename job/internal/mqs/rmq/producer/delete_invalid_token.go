package producer

import (
	"context"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
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
	return &DeleteInvalidTokenTask{
		ctx:      ctx,
		svcCtx:   svcCtx,
		producer: svcCtx.Config.ProducerConf.NewProducer(),
		cron:     gocron.NewScheduler(time.UTC),
	}
}

func (l *DeleteInvalidTokenTask) Start() {
	logx.Info("DeleteInvalidTokenTask producer start")
	err := l.producer.Start()
	logx.Must(err)

	// delete invalid token every 1 minute
	l.cron.Every(1).Minute().Do(func() {
		msg := &primitive.Message{
			Topic: "delete-invalid-token",
			Body:  []byte("all"),
		}
		res, err := l.producer.SendSync(context.Background(), msg)

		if err != nil {
			logx.Errorf("DeleteInvalidTokenTask send message error: %s\n", err.Error())
			return
		} else {
			logx.Infof("DeleteInvalidTokenTask send message success: %s\n", res.String())
		}
	})
}

func (l *DeleteInvalidTokenTask) Stop() {
	err := l.producer.Shutdown()
	if err != nil {
		logx.Errorw("DeleteInvalidTokenTask producer cannot shut down")
		return
	}
	l.cron.Stop()
}
