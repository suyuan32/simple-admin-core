package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/chimerakang/simple-admin-common/orm/ent/mixins"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("Role name | 角色名"),
		field.String("code").
			Comment("Role code for permission control in front end | 角色码，用于前端权限控制"),
		field.String("remark").Default("").
			Comment("Remark | 备注"),
		field.Uint32("sort").Default(0).
			Comment("Order number | 排序编号"),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menus", Menu.Type),
		edge.From("users", User.Type).Ref("roles"),
	}
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique(),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Role Table | 角色信息表"),
		entsql.Annotation{Table: "sys_roles"},
	}
}
