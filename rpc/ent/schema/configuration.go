package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Configuration holds the schema definition for the Configuration entity.
type Configuration struct {
	ent.Schema
}

// Fields of the Configuration.
func (Configuration) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Configurarion name | 配置名称"),
		field.String("key").Comment("Configuration key | 配置的键名"),
		field.String("value").Comment("Configuraion value | 配置的值"),
		field.String("category").Comment("Configuration category | 配置的分类"),
		field.String("remark").Optional().Comment("Remark | 备注"),
	}
}

// Edges of the Configuration.
func (Configuration) Edges() []ent.Edge {
	return nil
}

// Mixin of the Configuration.
func (Configuration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.SortMixin{},
		mixins.StateMixin{},
	}
}

// Indexes of the Configuration.
func (Configuration) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key"),
	}
}

// Annotations of the Configuration
func (Configuration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Dynamic Configuration Table | 动态配置表"),
		entsql.Annotation{Table: "sys_configuration"},
	}
}
