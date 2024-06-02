package base

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
)

// insert initial api data
func (l *InitDatabaseLogic) insertApiData() error {
	var apis []*ent.APICreate
	// USER
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/login").
		SetDescription("apiDesc.userLogin").
		SetAPIGroup("user").
		SetMethod("POST").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/register").
		SetDescription("apiDesc.userRegister").
		SetAPIGroup("user").
		SetMethod("POST").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/create").
		SetDescription("apiDesc.createUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/update").
		SetDescription("apiDesc.updateUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/change_password").
		SetDescription("apiDesc.userChangePassword").
		SetAPIGroup("user").
		SetMethod("POST").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/info").
		SetDescription("apiDesc.userInfo").
		SetAPIGroup("user").
		SetMethod("GET").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/list").
		SetDescription("apiDesc.userList").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/delete").
		SetDescription("apiDesc.deleteUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/perm").
		SetDescription("apiDesc.userPermissions").
		SetAPIGroup("user").
		SetMethod("GET").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/profile").
		SetDescription("apiDesc.userProfile").
		SetAPIGroup("user").
		SetMethod("GET").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/profile").
		SetDescription("apiDesc.updateProfile").
		SetAPIGroup("user").
		SetMethod("POST").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user/logout").
		SetDescription("apiDesc.logout").
		SetAPIGroup("user").
		SetMethod("GET").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/user").
		SetDescription("apiDesc.getUserById").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	// ROLE
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/role/create").
		SetDescription("apiDesc.createRole").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/role/update").
		SetDescription("apiDesc.updateRole").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/role/delete").
		SetDescription("apiDesc.deleteRole").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/role/list").
		SetDescription("apiDesc.roleList").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/role").
		SetDescription("apiDesc.getRoleById").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	// MENU

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu/create").
		SetDescription("apiDesc.createMenu").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu/update").
		SetDescription("apiDesc.updateMenu").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu/delete").
		SetDescription("apiDesc.deleteMenu").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu/list").
		SetDescription("apiDesc.menuList").
		SetAPIGroup("menu").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu/role/list").
		SetDescription("apiDesc.menuRoleList").
		SetAPIGroup("authority").
		SetMethod("GET").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu_param/create").
		SetDescription("apiDesc.createMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu_param/update").
		SetDescription("apiDesc.updateMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu_param/list").
		SetDescription("apiDesc.menuParamListByMenuId").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu_param/delete").
		SetDescription("apiDesc.deleteMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu_param").
		SetDescription("apiDesc.getMenuParamById").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/menu").
		SetDescription("apiDesc.getMenuById").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	// CAPTCHA

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/captcha").
		SetDescription("apiDesc.captcha").
		SetAPIGroup("captcha").
		SetMethod("GET").
		SetIsRequired(true),
	)

	// AUTHORIZATION

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/authority/api/create_or_update").
		SetDescription("apiDesc.createOrUpdateApiAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/authority/api/role").
		SetDescription("apiDesc.APIAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/authority/menu/create_or_update").
		SetDescription("apiDesc.createOrUpdateMenuAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/authority/menu/role").
		SetDescription("apiDesc.menuAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	// API

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/api/create").
		SetDescription("apiDesc.createApi").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/api/update").
		SetDescription("apiDesc.updateApi").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/api/delete").
		SetDescription("apiDesc.deleteAPI").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/api/list").
		SetDescription("apiDesc.APIList").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/api").
		SetDescription("apiDesc.getApiById").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	// DICTIONARY

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary").
		SetDescription("apiDesc.getDictionaryById").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary/create").
		SetDescription("apiDesc.createDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary/update").
		SetDescription("apiDesc.updateDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary/delete").
		SetDescription("apiDesc.deleteDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary_detail/delete").
		SetDescription("apiDesc.deleteDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary_detail").
		SetDescription("apiDesc.getDictionaryDetailById").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary_detail/create").
		SetDescription("apiDesc.createDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary_detail/update").
		SetDescription("apiDesc.updateDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary_detail/list").
		SetDescription("apiDesc.getDictionaryListDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dictionary/list").
		SetDescription("apiDesc.getDictionaryList").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/dict/:name").
		SetDescription("apiDesc.getDictionaryDetailByDictionaryName").
		SetAPIGroup("dictionary").
		SetMethod("GET"),
	)

	// OAUTH

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/oauth_provider/create").
		SetDescription("apiDesc.createProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/oauth_provider/update").
		SetDescription("apiDesc.updateProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/oauth_provider/delete").
		SetDescription("apiDesc.deleteProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/oauth_provider/list").
		SetDescription("apiDesc.getProviderList").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/oauth/login").
		SetDescription("apiDesc.oauthLogin").
		SetAPIGroup("oauth").
		SetMethod("POST").
		SetIsRequired(true),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/oauth_provider").
		SetDescription("apiDesc.getProviderById").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	// TOKEN

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/token/create").
		SetDescription("apiDesc.createToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/token/update").
		SetDescription("apiDesc.updateToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/token/delete").
		SetDescription("apiDesc.deleteToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/token/list").
		SetDescription("apiDesc.getTokenList").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/token/logout").
		SetDescription("apiDesc.forceLoggingOut").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/token").
		SetDescription("apiDesc.getTokenById").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	// DEPARTMENT

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/department/create").
		SetDescription("apiDesc.createDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/department/update").
		SetDescription("apiDesc.updateDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/department/delete").
		SetDescription("apiDesc.deleteDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/department/list").
		SetDescription("apiDesc.getDepartmentList").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/department").
		SetDescription("apiDesc.getDepartmentById").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	// POSITION

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/position/create").
		SetDescription("apiDesc.createPosition").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/position/update").
		SetDescription("apiDesc.updatePosition").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/position/delete").
		SetDescription("apiDesc.deletePosition").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/position/list").
		SetDescription("apiDesc.getPositionList").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Core").
		SetPath("/position").
		SetDescription("apiDesc.getPositionById").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	// TASK
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task/create").
		SetDescription("apiDesc.createTask").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task/update").
		SetDescription("apiDesc.updateTask").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task/delete").
		SetDescription("apiDesc.deleteTask").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task/list").
		SetDescription("apiDesc.getTaskList").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task").
		SetDescription("apiDesc.getTaskById").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	// TASK LOG
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task_log/create").
		SetDescription("apiDesc.createTaskLog").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task_log/update").
		SetDescription("apiDesc.updateTaskLog").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task_log/delete").
		SetDescription("apiDesc.deleteTaskLog").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task_log/list").
		SetDescription("apiDesc.getTaskLogList").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetServiceName("Job").
		SetPath("/task_log").
		SetDescription("apiDesc.getTaskLogById").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	err := l.svcCtx.DB.API.CreateBulk(apis...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}
