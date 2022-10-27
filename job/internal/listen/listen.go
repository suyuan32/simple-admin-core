package listen

import (
	"context"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/suyuan32/simple-admin-core/job/internal/config"
	"github.com/suyuan32/simple-admin-core/job/internal/svc"
)

func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	// add cron services
	services = append(services, Cron(c, ctx, svcContext)...)

	// add rocketmq services
	services = append(services, Rmq(c, ctx, svcContext)...)

	return services
}
