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
		core.NewDeleteInvalidTokenTask(ctx, svcCtx),
	}
}
