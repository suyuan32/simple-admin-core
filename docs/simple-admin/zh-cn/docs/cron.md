# 定时任务

> 添加定时任务

进入 job/internal/crons/core 添加任务，参考

```go
package core

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/job/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/model"
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
		err := l.svcCtx.DB.Where("status = ? and created_at < ?", 0, time.Now()).Delete(&model.Token{}).Error
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

	l.cron.StartAsync()
}

func (l *DeleteInvalidTokenTask) Stop() {
	logx.Info("cron DeleteInvalidTokenTask stop")
	l.cron.Stop()
}

```

### [Gocron](https://github.com/go-co-op/gocron)

>注意： 需要实现 Start 和 Stop 方法

> 添加任务监听

编辑 job/internal/listen/cron.go

```go
package listen

import (
	"context"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/suyuan32/simple-admin-core/job/internal/config"
	"github.com/suyuan32/simple-admin-core/job/internal/crons/core"
	"github.com/suyuan32/simple-admin-core/job/internal/svc"
)

// Cron service
func Cron(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		// add your cron task here
		core.NewDeleteInvalidTokenTask(ctx, svcCtx),
	}
}

```

> 启动服务

在 job 文件夹执行

```shell
go run core.go -f etc/core.yaml
```