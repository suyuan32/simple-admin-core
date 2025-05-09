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
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "path", Type: field.TypeString, Comment: "API path | API 路径"},
		{Name: "description", Type: field.TypeString, Comment: "API description | API 描述"},
		{Name: "api_group", Type: field.TypeString, Comment: "API group | API 分组"},
		{Name: "service_name", Type: field.TypeString, Comment: "Service name | 服务名称", Default: "Other"},
		{Name: "method", Type: field.TypeString, Comment: "HTTP method | HTTP 请求类型", Default: "POST"},
		{Name: "is_required", Type: field.TypeBool, Comment: "Whether is required | 是否必选", Default: false},
	}
	// SysApisTable holds the schema information for the "sys_apis" table.
	SysApisTable = &schema.Table{
		Name:       "sys_apis",
		Comment:    "API Table | API接口表",
		Columns:    SysApisColumns,
		PrimaryKey: []*schema.Column{SysApisColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_path_method",
				Unique:  true,
				Columns: []*schema.Column{SysApisColumns[3], SysApisColumns[7]},
			},
		},
	}
	// SysConfigurationColumns holds the columns for the "sys_configuration" table.
	SysConfigurationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "sort", Type: field.TypeUint32, Comment: "Sort Number | 排序编号", Default: 1},
		{Name: "state", Type: field.TypeBool, Nullable: true, Comment: "State true: normal false: ban | 状态 true 正常 false 禁用", Default: true},
		{Name: "name", Type: field.TypeString, Comment: "Configurarion name | 配置名称"},
		{Name: "key", Type: field.TypeString, Comment: "Configuration key | 配置的键名"},
		{Name: "value", Type: field.TypeString, Comment: "Configuraion value | 配置的值"},
		{Name: "category", Type: field.TypeString, Comment: "Configuration category | 配置的分类"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "Remark | 备注"},
	}
	// SysConfigurationTable holds the schema information for the "sys_configuration" table.
	SysConfigurationTable = &schema.Table{
		Name:       "sys_configuration",
		Comment:    "Dynamic Configuration Table | 动态配置表",
		Columns:    SysConfigurationColumns,
		PrimaryKey: []*schema.Column{SysConfigurationColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "configuration_key",
				Unique:  false,
				Columns: []*schema.Column{SysConfigurationColumns[6]},
			},
		},
	}
	// SysDepartmentsColumns holds the columns for the "sys_departments" table.
	SysDepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "sort", Type: field.TypeUint32, Comment: "Sort Number | 排序编号", Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Department name | 部门名称"},
		{Name: "ancestors", Type: field.TypeString, Nullable: true, Comment: "Parents' IDs | 父级列表"},
		{Name: "leader", Type: field.TypeString, Nullable: true, Comment: "Department leader | 部门负责人"},
		{Name: "phone", Type: field.TypeString, Nullable: true, Comment: "Leader's phone number | 负责人电话"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "Leader's email | 部门负责人电子邮箱"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "Remark | 备注"},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true, Comment: "Parent department ID | 父级部门ID", Default: 0},
	}
	// SysDepartmentsTable holds the schema information for the "sys_departments" table.
	SysDepartmentsTable = &schema.Table{
		Name:       "sys_departments",
		Comment:    "Department Table | 部门表",
		Columns:    SysDepartmentsColumns,
		PrimaryKey: []*schema.Column{SysDepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_departments_sys_departments_children",
				Columns:    []*schema.Column{SysDepartmentsColumns[11]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysDictionariesColumns holds the columns for the "sys_dictionaries" table.
	SysDictionariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "title", Type: field.TypeString, Comment: "The title shown in the ui | 展示名称 （建议配合i18n）"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "The name of dictionary for search | 字典搜索名称"},
		{Name: "desc", Type: field.TypeString, Nullable: true, Comment: "The description of dictionary | 字典的描述"},
	}
	// SysDictionariesTable holds the schema information for the "sys_dictionaries" table.
	SysDictionariesTable = &schema.Table{
		Name:       "sys_dictionaries",
		Comment:    "Dictionary Table | 字典信息表",
		Columns:    SysDictionariesColumns,
		PrimaryKey: []*schema.Column{SysDictionariesColumns[0]},
	}
	// SysDictionaryDetailsColumns holds the columns for the "sys_dictionary_details" table.
	SysDictionaryDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "sort", Type: field.TypeUint32, Comment: "Sort Number | 排序编号", Default: 1},
		{Name: "title", Type: field.TypeString, Comment: "The title shown in the ui | 展示名称 （建议配合i18n）"},
		{Name: "key", Type: field.TypeString, Comment: "key | 键"},
		{Name: "value", Type: field.TypeString, Comment: "value | 值"},
		{Name: "dictionary_id", Type: field.TypeUint64, Nullable: true, Comment: "Dictionary ID | 字典ID"},
	}
	// SysDictionaryDetailsTable holds the schema information for the "sys_dictionary_details" table.
	SysDictionaryDetailsTable = &schema.Table{
		Name:       "sys_dictionary_details",
		Comment:    "Dictionary Key/Value Table | 字典键值表",
		Columns:    SysDictionaryDetailsColumns,
		PrimaryKey: []*schema.Column{SysDictionaryDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dictionary_details_sys_dictionaries_dictionary_details",
				Columns:    []*schema.Column{SysDictionaryDetailsColumns[8]},
				RefColumns: []*schema.Column{SysDictionariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "sort", Type: field.TypeUint32, Comment: "Sort Number | 排序编号", Default: 1},
		{Name: "menu_level", Type: field.TypeUint32, Comment: "Menu level | 菜单层级"},
		{Name: "menu_type", Type: field.TypeUint32, Comment: "Menu type | 菜单类型 （菜单或目录）0 目录 1 菜单"},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "Index path | 菜单路由路径", Default: ""},
		{Name: "name", Type: field.TypeString, Comment: "Index name | 菜单名称"},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Comment: "Redirect path | 跳转路径 （外链）", Default: ""},
		{Name: "component", Type: field.TypeString, Nullable: true, Comment: "The path of vue file | 组件路径", Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Comment: "Disable status | 是否停用", Default: false},
		{Name: "service_name", Type: field.TypeString, Nullable: true, Comment: "Service Name | 服务名称", Default: "Other"},
		{Name: "permission", Type: field.TypeString, Nullable: true, Comment: "Permission symbol | 权限标识"},
		{Name: "title", Type: field.TypeString, Comment: "Menu name | 菜单显示标题"},
		{Name: "icon", Type: field.TypeString, Comment: "Menu icon | 菜单图标"},
		{Name: "hide_menu", Type: field.TypeBool, Nullable: true, Comment: "Hide menu | 是否隐藏菜单", Default: false},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Nullable: true, Comment: "Hide the breadcrumb | 隐藏面包屑", Default: false},
		{Name: "ignore_keep_alive", Type: field.TypeBool, Nullable: true, Comment: "Do not keep alive the tab | 取消页面缓存", Default: false},
		{Name: "hide_tab", Type: field.TypeBool, Nullable: true, Comment: "Hide the tab header | 隐藏页头", Default: false},
		{Name: "frame_src", Type: field.TypeString, Nullable: true, Comment: "Show iframe | 内嵌 iframe", Default: ""},
		{Name: "carry_param", Type: field.TypeBool, Nullable: true, Comment: "The route carries parameters or not | 携带参数", Default: false},
		{Name: "hide_children_in_menu", Type: field.TypeBool, Nullable: true, Comment: "Hide children menu or not | 隐藏所有子菜单", Default: false},
		{Name: "affix", Type: field.TypeBool, Nullable: true, Comment: "Affix tab | Tab 固定", Default: false},
		{Name: "dynamic_level", Type: field.TypeUint32, Nullable: true, Comment: "The maximum number of pages the router can open | 能打开的子TAB数", Default: 20},
		{Name: "real_path", Type: field.TypeString, Nullable: true, Comment: "The real path of the route without dynamic part | 菜单路由不包含参数部分", Default: ""},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true, Comment: "Parent menu ID | 父菜单ID", Default: 100000},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Comment:    "Menu Table | 菜单表",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[25]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "menu_name",
				Unique:  true,
				Columns: []*schema.Column{SysMenusColumns[7]},
			},
			{
				Name:    "menu_path",
				Unique:  true,
				Columns: []*schema.Column{SysMenusColumns[6]},
			},
		},
	}
	// SysOauthProvidersColumns holds the columns for the "sys_oauth_providers" table.
	SysOauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "The provider's name | 提供商名称"},
		{Name: "client_id", Type: field.TypeString, Comment: "The client id | 客户端 id"},
		{Name: "client_secret", Type: field.TypeString, Comment: "The client secret | 客户端密钥"},
		{Name: "redirect_url", Type: field.TypeString, Comment: "The redirect url | 跳转地址"},
		{Name: "scopes", Type: field.TypeString, Comment: "The scopes | 权限范围"},
		{Name: "auth_url", Type: field.TypeString, Comment: "The auth url of the provider | 认证地址"},
		{Name: "token_url", Type: field.TypeString, Comment: "The token url of the provider | 获取 token地址"},
		{Name: "auth_style", Type: field.TypeUint64, Comment: "The auth style, 0: auto detect 1: third party log in 2: log in with username and password | 鉴权方式 0 自动 1 第三方登录 2 使用用户名密码"},
		{Name: "info_url", Type: field.TypeString, Comment: "The URL to request user information by token | 用户信息请求地址"},
	}
	// SysOauthProvidersTable holds the schema information for the "sys_oauth_providers" table.
	SysOauthProvidersTable = &schema.Table{
		Name:       "sys_oauth_providers",
		Comment:    "Oauth Provider Configuration Table | 三方登录配置表",
		Columns:    SysOauthProvidersColumns,
		PrimaryKey: []*schema.Column{SysOauthProvidersColumns[0]},
	}
	// SysPositionsColumns holds the columns for the "sys_positions" table.
	SysPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "sort", Type: field.TypeUint32, Comment: "Sort Number | 排序编号", Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Position Name | 职位名称"},
		{Name: "code", Type: field.TypeString, Comment: "The code of position | 职位编码"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "Remark | 备注"},
	}
	// SysPositionsTable holds the schema information for the "sys_positions" table.
	SysPositionsTable = &schema.Table{
		Name:       "sys_positions",
		Comment:    "Position Table | 职位信息表",
		Columns:    SysPositionsColumns,
		PrimaryKey: []*schema.Column{SysPositionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "position_code",
				Unique:  true,
				Columns: []*schema.Column{SysPositionsColumns[6]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Role name | 角色名"},
		{Name: "code", Type: field.TypeString, Comment: "Role code for permission control in front end | 角色码，用于前端权限控制"},
		{Name: "default_router", Type: field.TypeString, Comment: "Default menu : dashboard | 默认登录页面", Default: "dashboard"},
		{Name: "remark", Type: field.TypeString, Comment: "Remark | 备注", Default: ""},
		{Name: "sort", Type: field.TypeUint32, Comment: "Order number | 排序编号", Default: 0},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Comment:    "Role Table | 角色信息表",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_code",
				Unique:  true,
				Columns: []*schema.Column{SysRolesColumns[5]},
			},
		},
	}
	// SysTokensColumns holds the columns for the "sys_tokens" table.
	SysTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Comment: "UUID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "uuid", Type: field.TypeUUID, Comment: " User's UUID | 用户的UUID"},
		{Name: "username", Type: field.TypeString, Comment: "Username | 用户名", Default: "unknown"},
		{Name: "token", Type: field.TypeString, Comment: "Token string | Token 字符串"},
		{Name: "source", Type: field.TypeString, Comment: "Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）"},
		{Name: "expired_at", Type: field.TypeTime, Comment: " Expire time | 过期时间"},
	}
	// SysTokensTable holds the schema information for the "sys_tokens" table.
	SysTokensTable = &schema.Table{
		Name:       "sys_tokens",
		Comment:    "Token Log Table | 令牌信息表",
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
				Columns: []*schema.Column{SysTokensColumns[8]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Comment: "UUID"},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "Delete Time | 删除日期"},
		{Name: "username", Type: field.TypeString, Unique: true, Comment: "User's login name | 登录名"},
		{Name: "password", Type: field.TypeString, Comment: "Password | 密码"},
		{Name: "nickname", Type: field.TypeString, Unique: true, Comment: "Nickname | 昵称"},
		{Name: "description", Type: field.TypeString, Nullable: true, Comment: "The description of user | 用户的描述信息"},
		{Name: "home_path", Type: field.TypeString, Comment: "The home page that the user enters after logging in | 用户登陆后进入的首页", Default: "/dashboard"},
		{Name: "mobile", Type: field.TypeString, Nullable: true, Comment: "Mobile number | 手机号"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "Email | 邮箱号"},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Comment: "Avatar | 头像路径", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Comment: "Department ID | 部门ID", Default: 1},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Comment:    "User Table | 用户信息表",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_users_sys_departments_departments",
				Columns:    []*schema.Column{SysUsersColumns[13]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[5], SysUsersColumns[11]},
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
	// UserPositionsColumns holds the columns for the "user_positions" table.
	UserPositionsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "position_id", Type: field.TypeUint64},
	}
	// UserPositionsTable holds the schema information for the "user_positions" table.
	UserPositionsTable = &schema.Table{
		Name:       "user_positions",
		Columns:    UserPositionsColumns,
		PrimaryKey: []*schema.Column{UserPositionsColumns[0], UserPositionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_positions_user_id",
				Columns:    []*schema.Column{UserPositionsColumns[0]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_positions_position_id",
				Columns:    []*schema.Column{UserPositionsColumns[1]},
				RefColumns: []*schema.Column{SysPositionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUint64},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysApisTable,
		SysConfigurationTable,
		SysDepartmentsTable,
		SysDictionariesTable,
		SysDictionaryDetailsTable,
		SysMenusTable,
		SysOauthProvidersTable,
		SysPositionsTable,
		SysRolesTable,
		SysTokensTable,
		SysUsersTable,
		RoleMenusTable,
		UserPositionsTable,
		UserRolesTable,
	}
)

func init() {
	SysApisTable.Annotation = &entsql.Annotation{
		Table: "sys_apis",
	}
	SysConfigurationTable.Annotation = &entsql.Annotation{
		Table: "sys_configuration",
	}
	SysDepartmentsTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysDepartmentsTable.Annotation = &entsql.Annotation{
		Table: "sys_departments",
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
	SysOauthProvidersTable.Annotation = &entsql.Annotation{
		Table: "sys_oauth_providers",
	}
	SysPositionsTable.Annotation = &entsql.Annotation{
		Table: "sys_positions",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	SysTokensTable.Annotation = &entsql.Annotation{
		Table: "sys_tokens",
	}
	SysUsersTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysUsersTable.Annotation = &entsql.Annotation{
		Table: "sys_users",
	}
	RoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
	UserPositionsTable.ForeignKeys[0].RefTable = SysUsersTable
	UserPositionsTable.ForeignKeys[1].RefTable = SysPositionsTable
	UserRolesTable.ForeignKeys[0].RefTable = SysUsersTable
	UserRolesTable.ForeignKeys[1].RefTable = SysRolesTable
}
