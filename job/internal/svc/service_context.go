package svc

import (
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/job/internal/config"
	"github.com/suyuan32/simple-admin-core/pkg/ent"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize database connection
	opts, err := c.DatabaseConf.NewEntOption(c.Redis)
	logx.Must(err)

	db := ent.NewClient(opts...)
	logx.Info("Initialize database connection successfully")

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
