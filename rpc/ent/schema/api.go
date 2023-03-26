package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type API struct {
	ent.Schema
}

func (API) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Comment("API path | API 路径"),
		field.String("description").Comment("API description | API 描述"),
		field.String("api_group").Comment("API group | API 分组"),
		field.String("method").Default("POST").Comment("HTTP method | HTTP 请求类型"),
	}
}

func (API) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (API) Edges() []ent.Edge {
	return nil
}

func (API) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path", "method").
			Unique(),
	}
}

func (API) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_apis"},
	}
}
