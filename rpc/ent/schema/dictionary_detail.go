package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type DictionaryDetail struct {
	ent.Schema
}

func (DictionaryDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("The title shown in the ui | 展示名称 （建议配合i18n）"),
		field.String("key").
			Comment("key | 键"),
		field.String("value").
			Comment("value | 值"),
		field.Uint64("dictionary_id").Optional().
			Comment("Dictionary ID | 字典ID"),
	}
}

func (DictionaryDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (DictionaryDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dictionaries", Dictionary.Type).Field("dictionary_id").Ref("dictionary_details").Unique(),
	}
}

func (DictionaryDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_dictionary_details"},
	}
}
