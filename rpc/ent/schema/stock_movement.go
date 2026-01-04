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

type StockMovement struct {
	ent.Schema
}

func (StockMovement) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("product_id", uuid.UUID{}).Comment("Product ID | 产品ID"),
		field.UUID("from_warehouse_id", uuid.UUID{}).Optional().Nillable().Comment("From Warehouse ID | 来源仓库ID"),
		field.UUID("to_warehouse_id", uuid.UUID{}).Optional().Nillable().Comment("To Warehouse ID | 目标仓库ID"),
		field.Int32("quantity").Comment("Quantity | 数量"),
		field.String("movement_type").Comment("Movement Type (IN/OUT/MOVE) | 移动类型"),
		field.String("reference").Comment("Reference | 关联单号"),
		field.String("details").Optional().Comment("Details | 详情"),
	}
}

func (StockMovement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (StockMovement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).Unique().Field("product_id").Required(),
		edge.To("from_warehouse", Warehouse.Type).Unique().Field("from_warehouse_id"),
		edge.To("to_warehouse", Warehouse.Type).Unique().Field("to_warehouse_id"),
	}
}

func (StockMovement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Stock Movement Table | 库存移动表"),
		entsql.Annotation{Table: "sys_stock_movements"},
	}
}
