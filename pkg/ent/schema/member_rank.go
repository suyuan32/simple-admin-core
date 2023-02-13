package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

type MemberRank struct {
	ent.Schema
}

func (MemberRank) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Rank name | 等级名称"),
		field.String("code").Comment("Rank code | 等级码"),
		field.String("description").Comment("Rank description | 等级描述"),
		field.String("remark").Comment("Remark | 备注"),
	}
}

func (MemberRank) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberRank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", Member.Type).Ref("ranks"),
	}
}

func (MemberRank) Indexes() []ent.Index {
	return nil
}

func (MemberRank) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "core_mms_rank"},
	}
}
