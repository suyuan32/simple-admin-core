package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
	mixins2 "github.com/suyuan32/simple-admin-core/rpc/ent/schema/mixins"
)

type Inventory struct {
	ent.Schema
}

func (Inventory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("product_id", uuid.UUID{}).Comment("Product ID | 产品ID"),
		field.UUID("warehouse_id", uuid.UUID{}).Comment("Warehouse ID | 仓库ID"),
		field.Int32("quantity").Min(0).Comment("Quantity | 数量"),
	}
}

func (Inventory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Inventory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).Unique().Field("product_id").Required(),
		edge.To("warehouse", Warehouse.Type).Unique().Field("warehouse_id").Required(),
	}
}

func (Inventory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Inventory Table | 库存表"),
		entsql.Annotation{Table: "sys_inventories"},
	}
}
