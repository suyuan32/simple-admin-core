package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Menu struct {
	ent.Schema
}

func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("parent_id").Default(100000).Optional().
			Comment("Parent menu ID | 父菜单ID").
			Annotations(entsql.WithComments(true)),
		field.Uint32("menu_level").
			Comment("Menu level | 菜单层级").
			Annotations(entsql.WithComments(true)),
		field.Uint32("menu_type").
			Comment("Menu type | 菜单类型 （菜单或目录）0 目录 1 菜单").
			Annotations(entsql.WithComments(true)),
		field.String("path").Optional().Default("").
			Comment("Index path | 菜单路由路径").
			Annotations(entsql.WithComments(true)),
		field.String("name").
			Comment("Index name | 菜单名称").
			Annotations(entsql.WithComments(true)),
		field.String("redirect").Optional().Default("").
			Comment("Redirect path | 跳转路径 （外链）").
			Annotations(entsql.WithComments(true)),
		field.String("component").Optional().Default("").
			Comment("The path of vue file | 组件路径").
			Annotations(entsql.WithComments(true)),
		field.Bool("disabled").Optional().Default(false).
			Comment("Disable status | 是否停用").
			Annotations(entsql.WithComments(true)),
		// meta
		field.String("title").
			Comment("Menu name | 菜单显示标题").
			Annotations(entsql.WithComments(true)),
		field.String("icon").
			Comment("Menu icon | 菜单图标").
			Annotations(entsql.WithComments(true)),
		field.Bool("hide_menu").Optional().Default(false).
			Comment("Hide menu | 是否隐藏菜单").
			Annotations(entsql.WithComments(true)),
		field.Bool("hide_breadcrumb").Optional().Default(false).
			Comment("Hide the breadcrumb | 隐藏面包屑").
			Annotations(entsql.WithComments(true)),
		field.Bool("ignore_keep_alive").Optional().Default(false).
			Comment("Do not keep alive the tab | 取消页面缓存").
			Annotations(entsql.WithComments(true)),
		field.Bool("hide_tab").Optional().Default(false).
			Comment("Hide the tab header | 隐藏页头").
			Annotations(entsql.WithComments(true)),
		field.String("frame_src").Optional().Default("").
			Comment("Show iframe | 内嵌 iframe").
			Annotations(entsql.WithComments(true)),
		field.Bool("carry_param").Optional().Default(false).
			Comment("The route carries parameters or not | 携带参数").
			Annotations(entsql.WithComments(true)),
		field.Bool("hide_children_in_menu").Optional().Default(false).
			Comment("Hide children menu or not | 隐藏所有子菜单").
			Annotations(entsql.WithComments(true)),
		field.Bool("affix").Optional().Default(false).
			Comment("Affix tab | Tab 固定").
			Annotations(entsql.WithComments(true)),
		field.Uint32("dynamic_level").Optional().Default(20).
			Comment("The maximum number of pages the router can open | 能打开的子TAB数").
			Annotations(entsql.WithComments(true)),
		field.String("real_path").Optional().Default("").
			Comment("The real path of the route without dynamic part | 菜单路由不包含参数部分").
			Annotations(entsql.WithComments(true)),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.SortMixin{},
	}
}

func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("path").Unique(),
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("menus"),
		edge.To("children", Menu.Type).From("parent").Unique().Field("parent_id"),
	}
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menus"},
	}
}
