# 全局变量

> 我们将全局变量定义到 internal/svc/servicecontext.go 中，并初始化 \
例如

```go
package svc

import (
	"github.com/suyuan32/simple-admin-core/api/internal/config"
	"github.com/suyuan32/simple-admin-core/api/internal/middleware"
	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	Redis     *redis.Redis
	Casbin    *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// initialize redis
	rds := c.RedisConf.NewRedis()
	logx.Info("Initialize redis connection successfully")

	// initialize database connection
	db, err := c.DB.NewGORM()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		return nil
	}
	logx.Info("Initialize database connection successful")

	// initialize casbin
	cbn := utils.NewCasbin(db)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:     rds,
		Casbin:    cbn,
	}
}

```

> 使用方法

```go
package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiLogic {
	return &CreateOrUpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateApiLogic) CreateOrUpdateApi(req *types.CreateOrUpdateApiReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateApi(context.Background(),
		&core.ApiInfo{
			Id:          req.Id,
			CreateAt:    req.CreateAt,
			Path:        req.Path,
			Description: req.Description,
			Group:       req.Group,
			Method:      req.Method,
		})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}

```

> 我们可以通过 l.svcCtx 访问全局变量