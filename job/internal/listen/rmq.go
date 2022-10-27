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
