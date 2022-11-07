package core

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/job/internal/svc"
	"github.com/suyuan32/simple-admin-core/pkg/ent/token"
	"github.com/suyuan32/simple-admin-core/pkg/enum"
)

type DeleteInvalidTokenTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	cron   *gocron.Scheduler
}

func NewDeleteInvalidTokenTask(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInvalidTokenTask {
	return &DeleteInvalidTokenTask{
		ctx:    ctx,
		svcCtx: svcCtx,
		cron:   gocron.NewScheduler(time.UTC),
	}
}

func (l *DeleteInvalidTokenTask) Start() {
	logx.Info("cron DeleteInvalidTokenTask start")

	// delete invalid token every 1 minute
	_, err := l.cron.Every(1).Minute().Do(func() {
		_, err := l.svcCtx.DB.Token.Delete().Where(token.And(token.StatusEQ(enum.StatusBanned),
			token.ExpiredAtLT(time.Now()))).Exec(l.ctx)
		if err != nil {
			logx.Errorf("DeleteInvalidTokenTask error: %s\n", err.Error())
		}
		logx.Info("successfully do the cron DeleteInvalidTokenTask")
		fmt.Println("successfully do the cron DeleteInvalidTokenTask")
	})

	if err != nil {
		logx.Error("cron error: %s\n", err.Error())
		return
	}

	l.cron.StartBlocking()
}

func (l *DeleteInvalidTokenTask) Stop() {
	logx.Info("cron DeleteInvalidTokenTask stop")
	l.cron.Stop()
}
