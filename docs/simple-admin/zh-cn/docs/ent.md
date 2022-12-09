## Ent 快速入门

#### [官方文档](https://entgo.io/zh/docs/getting-started/)
#### [schema中文文档(推荐)](https://suyuan32.github.io/ent-chinese-doc/#/zh-cn/getting-started)

## 实战
> 安装

```shell
go get -d entgo.io/ent/cmd/ent
```

> 初始化代码

在 pkg 执行

```shell
# 创建 User 模板
go run -mod=mod entgo.io/ent/cmd/ent init User

# 生成代码，使用 template ， simple admin core 添加了 Page 模板实现简便的分页查询
go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./ent/template/*.tmpl" ./ent/schema
```
> 定义数据模型

在 pkg/ent 中，一般只需要关注 schema 文件夹，里面定义了模型文件，其他文件夹和文件基本都是自动生成的, mixin 用于共用字段, 例如

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
		// field 定义字段， string 声明类型, comment 声明注释， default 声明默认值, unique 声明唯一
		field.String("value").Unique().Comment("role value for permission control in front end | 角色值，用于前端权限控制"),
		field.String("default_router").Default("dashboard").Comment("default menu : dashboard | 默认登录页面"),
		field.String("remark").Default("").Comment("remark | 备注"),
		field.Uint32("order_no").Default(0).Comment("order number | 排序编号"),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// 嵌入公用字段
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

> 初始化并添加全局引用

参考 rpc/internal/svc/service_context.go
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

注意： ent driver 有两种驱动，带缓存和不带缓存

> 带缓存 （会导致更新数据需要等待缓存时间过去才能看到更新，适合更新少的系统）

```go
db := ent.NewClient(
    ent.Log(logx.Info), // logger
    ent.Driver(c.DatabaseConf.GetCacheDriver(c.RedisConf)),
    ent.Debug(), // debug mode
)
```

> 不带缓存 (数据立即更新)

```go
db := ent.NewClient(
    ent.Log(logx.Info), // logger
    ent.Driver(c.DatabaseConf.GetNoCacheDriver()),
    ent.Debug(), // debug mode
)
```

> 使用效果

更新角色状态 rpc/internal/logic/update_role_status_logic.go

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

> 查询数据

查看文档 [断言](http://ent.ryansu.pro/#/zh-cn/predicates)

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

> 执行raw sql

若要支持纯 sql ，需要修改 makefile 生成代码， 添加 --feature sql/execquery

```shell
go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./pkg/ent/template/*.tmpl" ./pkg/ent/schema --feature sql/execquery
```

即可通过client.QueryContext 调用

```go
students, err := client.QueryContext(context.Background(), "select * from student")
```

> 项目默认添加了 page 模板

位于 ent/template/pagination.tmpl，生成代码时通过 --template glob="./pkg/ent/template/*.tmpl" 导入, 提供简便的分页功能,
如果你的其他项目也想要这个分页功能需要将 template 文件夹复制到新项目的ent文件夹中。

```go
apis, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
```

> 常见结果返回函数，用于query末尾

```go
// .ExecX() 只执行，不返回错误和数据
client.Student.UpdateOneID(1).SetName("Jack").ExecX(context.Background())

// .Exec() 执行并返回错误
err := client.Student.UpdateOneID(1).SetName("Jack").Exec(context.Background())

// .Save() 执行并返回结果数据和错误， 例如下面 s 保存 student 对象
s, err := client.Student.Create().
SetName("Jack").
SetAddress("Road").
SetAge(10).
Save(context.Background())

// .SaveX() 执行并返回结果数据， 例如下面 s 保存 student 对象
s := client.Student.Create().
SetName("Jack").
SetAddress("Road").
SetAge(10).
SaveX(context.Background())
```

