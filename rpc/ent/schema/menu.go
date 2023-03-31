package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Menu struct {
	ent.Schema
}

func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("parent_id").Default(100000).Optional().Comment("parent menu ID | 父菜单ID"),
		field.Uint32("menu_level").Comment("menu level | 菜单层级"),
		field.Uint32("menu_type").Comment("menu type | 菜单类型 （菜单或目录）0 目录 1 菜单"),
		field.String("path").Optional().Default("").Comment("index path | 菜单路由路径"),
		field.String("name").Comment("index name | 菜单名称"),
		field.String("redirect").Optional().Default("").Comment("redirect path | 跳转路径 （外链）"),
		field.String("component").Optional().Default("").Comment("the path of vue file | 组件路径"),
		field.Bool("disabled").Optional().Default(false).Comment("disable status | 是否停用"),
		// meta
		field.String("title").Comment("menu name | 菜单显示标题"),
		field.String("icon").Comment("menu icon | 菜单图标"),
		field.Bool("hide_menu").Optional().Default(false).Comment("hide menu | 是否隐藏菜单"),
		field.Bool("hide_breadcrumb").Optional().Default(false).Comment("hide the breadcrumb | 隐藏面包屑"),
		field.Bool("ignore_keep_alive").Optional().Default(false).Comment("do not keep alive the tab | 取消页面缓存"),
		field.Bool("hide_tab").Optional().Default(false).Comment("hide the tab header | 隐藏页头"),
		field.String("frame_src").Optional().Default("").Comment("show iframe | 内嵌 iframe"),
		field.Bool("carry_param").Optional().Default(false).Comment("the route carries parameters or not | 携带参数"),
		field.Bool("hide_children_in_menu").Optional().Default(false).Comment("hide children menu or not | 隐藏所有子菜单"),
		field.Bool("affix").Optional().Default(false).Comment("affix tab | Tab 固定"),
		field.Uint32("dynamic_level").Optional().Default(20).Comment("the maximum number of pages the router can open | 能打开的子TAB数"),
		field.String("real_path").Optional().Default("").Comment("the real path of the route without dynamic part | 菜单路由不包含参数部分"),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		mixins.SortMixin{},
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
