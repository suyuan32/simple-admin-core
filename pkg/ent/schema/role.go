package schema

import (
	"entgo.io/ent"
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
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menus", Menu.Type),
	}
}
