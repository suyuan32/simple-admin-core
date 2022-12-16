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
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.GetCacheDriver(c.RedisConf)),
		ent.Debug(), // debug mode
	)

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

Notice： There are two drivers for ent，cache and no cache.

> Cache （Will cause changes show slowly, suitable for system whose data have less changes）

```go
db := ent.NewClient(
    ent.Log(logx.Info), // logger
    ent.Driver(c.DatabaseConf.NewCacheDriver(c.RedisConf)),
    ent.Debug(), // debug mode
)
```

> No cache (Changes will show immediately)

```go
db := ent.NewClient(
    ent.Log(logx.Info), // logger
    ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
    ent.Debug(), // debug mode
)
```

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
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	// update redis

	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", role.ID), role.Name)
	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", role.ID), role.Value)
	err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", role.ID), strconv.Itoa(int(role.Status)))
	if err != nil {
		return nil, statuserr.NewInternalError(i18n.RedisError)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}

```

> Query data

See [Predicates](http://ent.ryansu.pro/#/zh-cn/predicates)

```go
package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/api"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListLogic) GetApiList(in *core.ApiPageReq) (*core.ApiListResp, error) {
	var predicates []predicate.API

	if in.Path != "" {
		predicates = append(predicates, api.PathContains(in.Path))
	}

	if in.Description != "" {
		predicates = append(predicates, api.DescriptionContains(in.Description))
	}

	if in.Method != "" {
		predicates = append(predicates, api.MethodContains(in.Method))
	}

	if in.Group != "" {
		predicates = append(predicates, api.APIGroupContains(in.Group))
	}

	apis, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.ApiListResp{}
	resp.Total = apis.PageDetails.Total

	for _, v := range apis.List {
		resp.Data = append(resp.Data, &core.ApiInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			Path:        v.Path,
			Description: v.Description,
			Group:       v.APIGroup,
			Method:      v.Method,
		})
	}

	return resp, nil
}

```

> query raw sql

If you want to execute raw sql ，you need to modify makefile ， add flag --feature sql/execquery

```shell
go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./pkg/ent/template/*.tmpl" ./pkg/ent/schema --feature sql/execquery
```

and then you can use client.QueryContext to execute raw sql

```go
students, err := client.QueryContext(context.Background(), "select * from student")
```

>  Project add pagination template by default, you can copy this template to other project

In ent/template/pagination.tmpl，add flag --template glob="./pkg/ent/template/*.tmpl" when generating code and 
then you can use pagination like below:

```go
apis, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
```

> Common functions used in query

```go
// .ExecX() execute query，don't return anything
client.Student.UpdateOneID(1).SetName("Jack").ExecX(context.Background())

// .Exec() execute query and return error
err := client.Student.UpdateOneID(1).SetName("Jack").Exec(context.Background())

// .Save() execute query and  return error and student data
s, err := client.Student.Create().
SetName("Jack").
SetAddress("Road").
SetAge(10).
Save(context.Background())

// .SaveX() execute query and  return student data without error
s := client.Student.Create().
SetName("Jack").
SetAddress("Road").
SetAge(10).
SaveX(context.Background())
```

> Ent schema import tool [ent import](https://github.com/ariga/entimport)