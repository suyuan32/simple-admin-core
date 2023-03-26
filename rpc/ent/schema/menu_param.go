package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type MenuParam struct {
	ent.Schema
}

func (MenuParam) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("Pass parameters via params or query | 参数类型"),
		field.String("key").Comment("The key of parameters | 参数键"),
		field.String("value").Comment("The value of parameters | 参数值"),
		field.Uint64("menu_id").Optional().Comment("The parent menu ID | 父级菜单ID"),
	}
}

func (MenuParam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (MenuParam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menus", Menu.Type).Field("menu_id").
			Ref("params").Unique(),
	}
}

func (MenuParam) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menu_params"},
	}
}
