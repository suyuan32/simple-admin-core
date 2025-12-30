package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
	mixins2 "github.com/suyuan32/simple-admin-core/rpc/ent/schema/mixins"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Product Name | 产品名称"),
		field.String("sku").Unique().Comment("SKU | 库存单位"),
		field.String("description").Optional().Comment("Description | 描述"),
		field.Float("price").Min(0).Comment("Price | 价格"),
		field.String("unit").Comment("Unit | 单位"),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Product Table | 产品表"),
		entsql.Annotation{Table: "sys_products"},
	}
}
