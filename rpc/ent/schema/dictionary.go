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
		field.String("title").Comment("the title shown in the ui | 展示名称 （建议配合i18n）"),
		field.String("name").Unique().Comment("the name of dictionary for search | 字典搜索名称"),
		field.String("desc").Comment("the status of dictionary (true enable | false disable) | 字典状态"),
	}
}

func (Dictionary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
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
		entsql.Annotation{Table: "sys_dictionaries"},
	}
}
