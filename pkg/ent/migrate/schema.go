// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysApisColumns holds the columns for the "sys_apis" table.
	SysApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "path", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "api_group", Type: field.TypeString},
		{Name: "method", Type: field.TypeString, Default: "POST"},
	}
	// SysApisTable holds the schema information for the "sys_apis" table.
	SysApisTable = &schema.Table{
		Name:       "sys_apis",
		Columns:    SysApisColumns,
		PrimaryKey: []*schema.Column{SysApisColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_path_method",
				Unique:  true,
				Columns: []*schema.Column{SysApisColumns[3], SysApisColumns[6]},
			},
		},
	}
	// SysDictionariesColumns holds the columns for the "sys_dictionaries" table.
	SysDictionariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "title", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "desc", Type: field.TypeString},
	}
	// SysDictionariesTable holds the schema information for the "sys_dictionaries" table.
	SysDictionariesTable = &schema.Table{
		Name:       "sys_dictionaries",
		Columns:    SysDictionariesColumns,
		PrimaryKey: []*schema.Column{SysDictionariesColumns[0]},
	}
	// SysDictionaryDetailsColumns holds the columns for the "sys_dictionary_details" table.
	SysDictionaryDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "title", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "dictionary_dictionary_details", Type: field.TypeUint64, Nullable: true},
	}
	// SysDictionaryDetailsTable holds the schema information for the "sys_dictionary_details" table.
	SysDictionaryDetailsTable = &schema.Table{
		Name:       "sys_dictionary_details",
		Columns:    SysDictionaryDetailsColumns,
		PrimaryKey: []*schema.Column{SysDictionaryDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dictionary_details_sys_dictionaries_dictionary_details",
				Columns:    []*schema.Column{SysDictionaryDetailsColumns[7]},
				RefColumns: []*schema.Column{SysDictionariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "menu_level", Type: field.TypeUint32},
		{Name: "menu_type", Type: field.TypeUint32},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "component", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "order_no", Type: field.TypeUint32, Default: 0},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "title", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
		{Name: "hide_menu", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "current_active_menu", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "ignore_keep_alive", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "hide_tab", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "frame_src", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "carry_param", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "hide_children_in_menu", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "affix", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "dynamic_level", Type: field.TypeUint32, Nullable: true, Default: 20},
		{Name: "real_path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[24]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenuParamsColumns holds the columns for the "sys_menu_params" table.
	SysMenuParamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "menu_params", Type: field.TypeUint64, Nullable: true},
	}
	// SysMenuParamsTable holds the schema information for the "sys_menu_params" table.
	SysMenuParamsTable = &schema.Table{
		Name:       "sys_menu_params",
		Columns:    SysMenuParamsColumns,
		PrimaryKey: []*schema.Column{SysMenuParamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menu_params_sys_menus_params",
				Columns:    []*schema.Column{SysMenuParamsColumns[6]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysOauthProvidersColumns holds the columns for the "sys_oauth_providers" table.
	SysOauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "client_id", Type: field.TypeString},
		{Name: "client_secret", Type: field.TypeString},
		{Name: "redirect_url", Type: field.TypeString},
		{Name: "scopes", Type: field.TypeString},
		{Name: "auth_url", Type: field.TypeString},
		{Name: "token_url", Type: field.TypeString},
		{Name: "auth_style", Type: field.TypeUint64},
		{Name: "info_url", Type: field.TypeString},
	}
	// SysOauthProvidersTable holds the schema information for the "sys_oauth_providers" table.
	SysOauthProvidersTable = &schema.Table{
		Name:       "sys_oauth_providers",
		Columns:    SysOauthProvidersColumns,
		PrimaryKey: []*schema.Column{SysOauthProvidersColumns[0]},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString, Unique: true},
		{Name: "default_router", Type: field.TypeString, Default: "dashboard"},
		{Name: "remark", Type: field.TypeString, Default: ""},
		{Name: "order_no", Type: field.TypeUint32, Default: 0},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
	}
	// SysTenantColumns holds the columns for the "sys_tenant" table.
	SysTenantColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "uuid", Type: field.TypeString, Default: "70a26f02-4da8-4c31-aedf-8708c3cb42e5"},
		{Name: "level", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "contact", Type: field.TypeString, Nullable: true},
		{Name: "mobile", Type: field.TypeString, Nullable: true},
		{Name: "sort_no", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "pid", Type: field.TypeUint64, Nullable: true},
	}
	// SysTenantTable holds the schema information for the "sys_tenant" table.
	SysTenantTable = &schema.Table{
		Name:       "sys_tenant",
		Columns:    SysTenantColumns,
		PrimaryKey: []*schema.Column{SysTenantColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_tenant_sys_tenant_children",
				Columns:    []*schema.Column{SysTenantColumns[13]},
				RefColumns: []*schema.Column{SysTenantColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tenant_uuid_pid",
				Unique:  true,
				Columns: []*schema.Column{SysTenantColumns[4], SysTenantColumns[13]},
			},
			{
				Name:    "tenant_sort_no",
				Unique:  false,
				Columns: []*schema.Column{SysTenantColumns[12]},
			},
		},
	}
	// SysTokensColumns holds the columns for the "sys_tokens" table.
	SysTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "uuid", Type: field.TypeString},
		{Name: "token", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "source", Type: field.TypeString},
		{Name: "expired_at", Type: field.TypeTime},
	}
	// SysTokensTable holds the schema information for the "sys_tokens" table.
	SysTokensTable = &schema.Table{
		Name:       "sys_tokens",
		Columns:    SysTokensColumns,
		PrimaryKey: []*schema.Column{SysTokensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "token_uuid",
				Unique:  false,
				Columns: []*schema.Column{SysTokensColumns[4]},
			},
			{
				Name:    "token_expired_at",
				Unique:  false,
				Columns: []*schema.Column{SysTokensColumns[7]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "uuid", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "side_mode", Type: field.TypeString, Nullable: true, Default: "dark"},
		{Name: "base_color", Type: field.TypeString, Nullable: true, Default: "#fff"},
		{Name: "active_color", Type: field.TypeString, Nullable: true, Default: "#1890ff"},
		{Name: "role_id", Type: field.TypeUint64, Nullable: true, Default: 2},
		{Name: "mobile", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Default: "", SchemaType: map[string]string{"mysql": "varchar(512)"}},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[5], SysUsersColumns[13]},
			},
		},
	}
	// RoleMenusColumns holds the columns for the "role_menus" table.
	RoleMenusColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUint64},
		{Name: "menu_id", Type: field.TypeUint64},
	}
	// RoleMenusTable holds the schema information for the "role_menus" table.
	RoleMenusTable = &schema.Table{
		Name:       "role_menus",
		Columns:    RoleMenusColumns,
		PrimaryKey: []*schema.Column{RoleMenusColumns[0], RoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_menus_role_id",
				Columns:    []*schema.Column{RoleMenusColumns[0]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_menus_menu_id",
				Columns:    []*schema.Column{RoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TenantUsersColumns holds the columns for the "tenant_users" table.
	TenantUsersColumns = []*schema.Column{
		{Name: "tenant_id", Type: field.TypeUint64},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// TenantUsersTable holds the schema information for the "tenant_users" table.
	TenantUsersTable = &schema.Table{
		Name:       "tenant_users",
		Columns:    TenantUsersColumns,
		PrimaryKey: []*schema.Column{TenantUsersColumns[0], TenantUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tenant_users_tenant_id",
				Columns:    []*schema.Column{TenantUsersColumns[0]},
				RefColumns: []*schema.Column{SysTenantColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tenant_users_user_id",
				Columns:    []*schema.Column{TenantUsersColumns[1]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysApisTable,
		SysDictionariesTable,
		SysDictionaryDetailsTable,
		SysMenusTable,
		SysMenuParamsTable,
		SysOauthProvidersTable,
		SysRolesTable,
		SysTenantTable,
		SysTokensTable,
		SysUsersTable,
		RoleMenusTable,
		TenantUsersTable,
	}
)

func init() {
	SysApisTable.Annotation = &entsql.Annotation{
		Table: "sys_apis",
	}
	SysDictionariesTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionaries",
	}
	SysDictionaryDetailsTable.ForeignKeys[0].RefTable = SysDictionariesTable
	SysDictionaryDetailsTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionary_details",
	}
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
	SysMenuParamsTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenuParamsTable.Annotation = &entsql.Annotation{
		Table: "sys_menu_params",
	}
	SysOauthProvidersTable.Annotation = &entsql.Annotation{
		Table: "sys_oauth_providers",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	SysTenantTable.ForeignKeys[0].RefTable = SysTenantTable
	SysTenantTable.Annotation = &entsql.Annotation{
		Table: "sys_tenant",
	}
	SysTokensTable.Annotation = &entsql.Annotation{
		Table: "sys_tokens",
	}
	SysUsersTable.Annotation = &entsql.Annotation{
		Table: "sys_users",
	}
	RoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
	TenantUsersTable.ForeignKeys[0].RefTable = SysTenantTable
	TenantUsersTable.ForeignKeys[1].RefTable = SysUsersTable
}
