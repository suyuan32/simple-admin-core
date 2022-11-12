## Ent 快速入门

#### [官方文档](https://entgo.io/zh/docs/getting-started/)
#### [schema中文文档](https://suyuan32.github.io/ent-chinese-doc/#/zh-cn/getting-started)

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

使用 c.DatabaseConf.NewEntOption(c.RedisConf) 初始化配置，使用 ent.NewClient(opts...) 创建 client, 即可
全局使用

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
			return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotFound)
		default:
			logx.Errorw(logmsg.DATABASE_ERROR, logx.Field("detail", err.Error()))
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
