package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.“
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Default(uuid.NewString()).Comment("tenant's UUID | 租户的UUID"),
		field.Uint64("pid").Optional().Comment("parent id | 父级ID"),
		field.Uint32("level").Comment("tenant's level | 租户级别（含部门）"),
		field.String("name").Unique().Comment("tenant's name | 租户的名称"),
		field.String("account").Unique().Comment("tenant's account | 租户登录账号"),
		field.Time("start_time").Comment("start_time | 租期的开始时间").
			Default(time.Now),
		field.Time("end_time").Optional().Comment("end_time ｜ 租期的结束时间"),
		field.String("contact").Optional().Comment("contact | 客户联系人"),
		field.String("mobile").Optional().Comment("mobile | 客户联系电话"),
		field.Uint32("sort_no").Optional().Default(0).Comment("sort number | 显示排序"),
	}
}

func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		// tenant contains users
		edge.To("users", User.Type),
		// tenant children and parent relations
		edge.To("children", Tenant.Type).From("parent").Unique().Field("pid"),
	}
}

func (Tenant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "pid").Unique(),
		index.Fields("sort_no"),
	}
}

func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_tenant"},
	}
}
