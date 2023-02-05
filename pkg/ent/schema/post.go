package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Post Name | 岗位名称"),
		field.String("code").Comment("The code of post | 岗位编码"),
		field.String("remark").Comment("Remark | 备注"),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("post"),
	}
}

func (Post) Indexes() []ent.Index {
	return nil
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_posts"},
	}
}
