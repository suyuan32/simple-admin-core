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

	services = append(services, Cron(c, ctx, svcContext)...)

	return services
}
