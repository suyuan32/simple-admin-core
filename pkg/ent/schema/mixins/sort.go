package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SortMixin implements the ent.Mixin for sharing
// sort fields with package schemas.
type SortMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (SortMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("sort").Default(1).Comment("Sort number | 排序编号"),
	}
}
