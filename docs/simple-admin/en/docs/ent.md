## Ent

#### [Official Doc](https://entgo.io/docs/getting-started)

## Quick Start

> Install ent

```shell
go get -d entgo.io/ent/cmd/ent
```

> Initialize code

Run in pkg directory

```shell
# Creat User schema
go run -mod=mod entgo.io/ent/cmd/ent init User

# Generate code，use template ， simple admin core add Page template to make it easier to do pagination
go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./ent/template/*.tmpl" ./ent/schema
```
> Defined schema

In pkg/ent ，you mainly modify files in schema directory. It is used to define all the models，most of other files in other directories are 
generated. Mixin is used to define common fields, such as

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("role name | 角色名"),
		field.String("value").Unique().Comment("role value for permission control in front end | 角色值，用于前端权限控制"),
		field.String("default_router").Default("dashboard").Comment("default menu : dashboard | 默认登录页面"),
		field.String("remark").Default("").Comment("remark | 备注"),
		field.Uint32("order_no").Default(0).Comment("order number | 排序编号"),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// common field
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		// 声明关系, ent 的关系用 edge 表示
		edge.To("menus", Menu.Type),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_roles"},
	}
}

```

> Initialize 

See rpc/internal/svc/service_context.go
```go
package svc

import (
	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	opts, err := c.DatabaseConf.NewEntOption(c.RedisConf)
	logx.Must(err)

	db := ent.NewClient(opts...)
	logx.Info("Initialize database connection successfully")

	// initialize redis
	rds := c.RedisConf.NewRedis()
	if !rds.Ping() {
		logx.Error("Initialize redis failed")
		return nil
	}
	logx.Info("Initialize redis connection successfully")

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
	}
}

```

Use c.DatabaseConf.NewEntOption(c.RedisConf) initialize options，use ent.NewClient(opts...) create client and then you
can use it globally

> Usage in logic

Update role status.  rpc/internal/logic/update_role_status_logic.go

```go
package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *core.StatusCodeReq) (*core.BaseResp, error) {
	role, err := l.svcCtx.DB.Role.UpdateOneID(in.Id).SetStatus(uint8(in.Status)).Save(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	// update redis

	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", role.ID), role.Name)
	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", role.ID), role.Value)
	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", role.ID), strconv.Itoa(int(role.Status)))
	if err != nil {
		return nil, statuserr.NewInternalError(errorx.RedisError)
	}

	return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
}

```
