package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Department struct {
	ent.Schema
}

func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("Department name | 部门名称").
			Annotations(entsql.WithComments(true)),
		field.String("ancestors").Optional().
			Comment("Parents' IDs | 父级列表").
			Annotations(entsql.WithComments(true)),
		field.String("leader").
			Comment("Department leader | 部门负责人").Optional().
			Annotations(entsql.WithComments(true)),
		field.String("phone").
			Comment("Leader's phone number | 负责人电话").Optional().
			Annotations(entsql.WithComments(true)),
		field.String("email").
			Comment("Leader's email | 部门负责人电子邮箱").Optional().
			Annotations(entsql.WithComments(true)),
		field.String("remark").Optional().
			Comment("Remark | 备注").
			Annotations(entsql.WithComments(true)),
		field.Uint64("parent_id").Optional().Default(0).
			Comment("Parent department ID | 父级部门ID").
			Annotations(entsql.WithComments(true)),
	}
}

func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Department.Type).From("parent").Unique().Field("parent_id"),
		edge.From("users", User.Type).Ref("departments"),
	}
}

func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_departments"},
	}
}
