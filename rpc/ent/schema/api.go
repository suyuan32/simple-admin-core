package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/chimerakang/simple-admin-common/orm/ent/mixins"
)

type API struct {
	ent.Schema
}

func (API) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").
			Comment("API path | API 路径"),
		field.String("description").
			Comment("API description | API 描述"),
		field.String("api_group").
			Comment("API group | API 分组"),
		field.String("service_name").
			Comment("Service name | 服务名称").
			Default("Other"),
		field.String("method").Default("POST").
			Comment("HTTP method | HTTP 请求类型"),
		field.Bool("is_required").Default(false).
			Comment("Whether is required | 是否必选"),
	}
}

func (API) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
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
		entsql.WithComments(true),
		schema.Comment("API Table | API接口表"),
		entsql.Annotation{Table: "sys_apis"},
	}
}
