// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"
)

// Menu Table | 菜单表
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Sort Number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
	// Parent menu ID | 父菜单ID
	ParentID uint64 `json:"parent_id,omitempty"`
	// Menu level | 菜单层级
	MenuLevel uint32 `json:"menu_level,omitempty"`
	// Menu type | 菜单类型 （菜单或目录）0 目录 1 菜单
	MenuType uint32 `json:"menu_type,omitempty"`
	// Index path | 菜单路由路径
	Path string `json:"path,omitempty"`
	// Index name | 菜单名称
	Name string `json:"name,omitempty"`
	// Redirect path | 跳转路径 （外链）
	Redirect string `json:"redirect,omitempty"`
	// The path of vue file | 组件路径
	Component string `json:"component,omitempty"`
	// Disable status | 是否停用
	Disabled bool `json:"disabled,omitempty"`
	// Service Name | 服务名称
	ServiceName string `json:"service_name,omitempty"`
	// Permission symbol | 权限标识
	Permission string `json:"permission,omitempty"`
	// Menu name | 菜单显示标题
	Title string `json:"title,omitempty"`
	// Menu icon | 菜单图标
	Icon string `json:"icon,omitempty"`
	// Hide menu | 是否隐藏菜单
	HideMenu bool `json:"hide_menu,omitempty"`
	// Hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hide_breadcrumb,omitempty"`
	// Do not keep alive the tab | 取消页面缓存
	IgnoreKeepAlive bool `json:"ignore_keep_alive,omitempty"`
	// Hide the tab header | 隐藏页头
	HideTab bool `json:"hide_tab,omitempty"`
	// Show iframe | 内嵌 iframe
	FrameSrc string `json:"frame_src,omitempty"`
	// The route carries parameters or not | 携带参数
	CarryParam bool `json:"carry_param,omitempty"`
	// Hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hide_children_in_menu,omitempty"`
	// Affix tab | Tab 固定
	Affix bool `json:"affix,omitempty"`
	// The maximum number of pages the router can open | 能打开的子TAB数
	DynamicLevel uint32 `json:"dynamic_level,omitempty"`
	// The real path of the route without dynamic part | 菜单路由不包含参数部分
	RealPath string `json:"real_path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Menu `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ChildrenOrErr() ([]*Menu, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldDisabled, menu.FieldHideMenu, menu.FieldHideBreadcrumb, menu.FieldIgnoreKeepAlive, menu.FieldHideTab, menu.FieldCarryParam, menu.FieldHideChildrenInMenu, menu.FieldAffix:
			values[i] = new(sql.NullBool)
		case menu.FieldID, menu.FieldSort, menu.FieldParentID, menu.FieldMenuLevel, menu.FieldMenuType, menu.FieldDynamicLevel:
			values[i] = new(sql.NullInt64)
		case menu.FieldPath, menu.FieldName, menu.FieldRedirect, menu.FieldComponent, menu.FieldServiceName, menu.FieldPermission, menu.FieldTitle, menu.FieldIcon, menu.FieldFrameSrc, menu.FieldRealPath:
			values[i] = new(sql.NullString)
		case menu.FieldCreatedAt, menu.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case menu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case menu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case menu.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				m.Sort = uint32(value.Int64)
			}
		case menu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				m.ParentID = uint64(value.Int64)
			}
		case menu.FieldMenuLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_level", values[i])
			} else if value.Valid {
				m.MenuLevel = uint32(value.Int64)
			}
		case menu.FieldMenuType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_type", values[i])
			} else if value.Valid {
				m.MenuType = uint32(value.Int64)
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				m.Redirect = value.String
			}
		case menu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				m.Component = value.String
			}
		case menu.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				m.Disabled = value.Bool
			}
		case menu.FieldServiceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_name", values[i])
			} else if value.Valid {
				m.ServiceName = value.String
			}
		case menu.FieldPermission:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permission", values[i])
			} else if value.Valid {
				m.Permission = value.String
			}
		case menu.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case menu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				m.Icon = value.String
			}
		case menu.FieldHideMenu:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_menu", values[i])
			} else if value.Valid {
				m.HideMenu = value.Bool
			}
		case menu.FieldHideBreadcrumb:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_breadcrumb", values[i])
			} else if value.Valid {
				m.HideBreadcrumb = value.Bool
			}
		case menu.FieldIgnoreKeepAlive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ignore_keep_alive", values[i])
			} else if value.Valid {
				m.IgnoreKeepAlive = value.Bool
			}
		case menu.FieldHideTab:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_tab", values[i])
			} else if value.Valid {
				m.HideTab = value.Bool
			}
		case menu.FieldFrameSrc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field frame_src", values[i])
			} else if value.Valid {
				m.FrameSrc = value.String
			}
		case menu.FieldCarryParam:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field carry_param", values[i])
			} else if value.Valid {
				m.CarryParam = value.Bool
			}
		case menu.FieldHideChildrenInMenu:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_children_in_menu", values[i])
			} else if value.Valid {
				m.HideChildrenInMenu = value.Bool
			}
		case menu.FieldAffix:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field affix", values[i])
			} else if value.Valid {
				m.Affix = value.Bool
			}
		case menu.FieldDynamicLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dynamic_level", values[i])
			} else if value.Valid {
				m.DynamicLevel = uint32(value.Int64)
			}
		case menu.FieldRealPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_path", values[i])
			} else if value.Valid {
				m.RealPath = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Menu entity.
func (m *Menu) QueryRoles() *RoleQuery {
	return NewMenuClient(m.config).QueryRoles(m)
}

// QueryParent queries the "parent" edge of the Menu entity.
func (m *Menu) QueryParent() *MenuQuery {
	return NewMenuClient(m.config).QueryParent(m)
}

// QueryChildren queries the "children" edge of the Menu entity.
func (m *Menu) QueryChildren() *MenuQuery {
	return NewMenuClient(m.config).QueryChildren(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", m.Sort))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", m.ParentID))
	builder.WriteString(", ")
	builder.WriteString("menu_level=")
	builder.WriteString(fmt.Sprintf("%v", m.MenuLevel))
	builder.WriteString(", ")
	builder.WriteString("menu_type=")
	builder.WriteString(fmt.Sprintf("%v", m.MenuType))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(m.Path)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("redirect=")
	builder.WriteString(m.Redirect)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(m.Component)
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", m.Disabled))
	builder.WriteString(", ")
	builder.WriteString("service_name=")
	builder.WriteString(m.ServiceName)
	builder.WriteString(", ")
	builder.WriteString("permission=")
	builder.WriteString(m.Permission)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(m.Title)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(m.Icon)
	builder.WriteString(", ")
	builder.WriteString("hide_menu=")
	builder.WriteString(fmt.Sprintf("%v", m.HideMenu))
	builder.WriteString(", ")
	builder.WriteString("hide_breadcrumb=")
	builder.WriteString(fmt.Sprintf("%v", m.HideBreadcrumb))
	builder.WriteString(", ")
	builder.WriteString("ignore_keep_alive=")
	builder.WriteString(fmt.Sprintf("%v", m.IgnoreKeepAlive))
	builder.WriteString(", ")
	builder.WriteString("hide_tab=")
	builder.WriteString(fmt.Sprintf("%v", m.HideTab))
	builder.WriteString(", ")
	builder.WriteString("frame_src=")
	builder.WriteString(m.FrameSrc)
	builder.WriteString(", ")
	builder.WriteString("carry_param=")
	builder.WriteString(fmt.Sprintf("%v", m.CarryParam))
	builder.WriteString(", ")
	builder.WriteString("hide_children_in_menu=")
	builder.WriteString(fmt.Sprintf("%v", m.HideChildrenInMenu))
	builder.WriteString(", ")
	builder.WriteString("affix=")
	builder.WriteString(fmt.Sprintf("%v", m.Affix))
	builder.WriteString(", ")
	builder.WriteString("dynamic_level=")
	builder.WriteString(fmt.Sprintf("%v", m.DynamicLevel))
	builder.WriteString(", ")
	builder.WriteString("real_path=")
	builder.WriteString(m.RealPath)
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
