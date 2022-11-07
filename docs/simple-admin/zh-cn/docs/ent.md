## Ent 快速入门

#### [官方文档](https://entgo.io/zh/docs/getting-started/)

## 实战
> 安装

```shell
go get -d entgo.io/ent/cmd/ent

# 创建 User 模板
go run -mod=mod entgo.io/ent/cmd/ent init User

# 生成代码
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


#### [schema中文参考（上）](https://blog.csdn.net/OpenSkill/article/details/108271774)
#### [schema （下）](https://blog.csdn.net/OpenSkill/article/details/108289545)

#### [代码生成及CRUD](https://blog.csdn.net/OpenSkill/article/details/108373737)