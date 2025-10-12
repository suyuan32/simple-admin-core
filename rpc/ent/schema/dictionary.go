package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Dictionary struct {
	ent.Schema
}

func (Dictionary) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("The title shown in the ui | 展示名称 （建议配合i18n）"),
		field.String("name").Unique().
			Comment("The name of dictionary for search | 字典搜索名称"),
		field.String("desc").
			Comment("The description of dictionary | 字典的描述").
			Optional(),
		field.Bool("is_public").Default(false).Comment("Whether to be public for everyone | 是否公开词典，无需登录即可访问"),
	}
}

func (Dictionary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

func (Dictionary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dictionary_details", DictionaryDetail.Type),
	}
}

func (Dictionary) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Dictionary Table | 字典信息表"),
		entsql.Annotation{Table: "sys_dictionaries"},
	}
}
