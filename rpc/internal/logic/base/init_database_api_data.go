package base

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
)

// insert init api data
func (l *InitDatabaseLogic) insertApiData() error {
	var apis []*ent.APICreate
	// USER
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/login").
		SetDescription("apiDesc.userLogin").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/register").
		SetDescription("apiDesc.userRegister").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/create").
		SetDescription("apiDesc.createUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/update").
		SetDescription("apiDesc.updateUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/change_password").
		SetDescription("apiDesc.userChangePassword").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/info").
		SetDescription("apiDesc.userInfo").
		SetAPIGroup("user").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/list").
		SetDescription("apiDesc.userList").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/delete").
		SetDescription("apiDesc.deleteUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/perm").
		SetDescription("apiDesc.userPermissions").
		SetAPIGroup("user").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/profile").
		SetDescription("apiDesc.userProfile").
		SetAPIGroup("user").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/profile").
		SetDescription("apiDesc.updateProfile").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/logout").
		SetDescription("apiDesc.logout").
		SetAPIGroup("user").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user").
		SetDescription("apiDesc.getUserById").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	// ROLE
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/role/create").
		SetDescription("apiDesc.createRole").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/role/update").
		SetDescription("apiDesc.updateRole").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/role/delete").
		SetDescription("apiDesc.deleteRole").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/role/list").
		SetDescription("apiDesc.roleList").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/role").
		SetDescription("apiDesc.getRoleById").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	// MENU

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/create").
		SetDescription("apiDesc.createMenu").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/update").
		SetDescription("apiDesc.updateMenu").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/delete").
		SetDescription("apiDesc.deleteMenu").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/list").
		SetDescription("apiDesc.menuList").
		SetAPIGroup("menu").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/role/list").
		SetDescription("apiDesc.menuRoleList").
		SetAPIGroup("authority").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu_param/create").
		SetDescription("apiDesc.createMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu_param/update").
		SetDescription("apiDesc.updateMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu_param/list").
		SetDescription("apiDesc.menuParamListByMenuId").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu_param/delete").
		SetDescription("apiDesc.deleteMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu_param").
		SetDescription("apiDesc.getMenuParamById").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu").
		SetDescription("apiDesc.getMenuById").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	// CAPTCHA

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/captcha").
		SetDescription("apiDesc.captcha").
		SetAPIGroup("captcha").
		SetMethod("GET"),
	)

	// AUTHORIZATION

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/authority/api/create_or_update").
		SetDescription("apiDesc.createOrUpdateApiAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/authority/api/role").
		SetDescription("apiDesc.APIAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/authority/menu/create_or_update").
		SetDescription("apiDesc.createOrUpdateMenuAuthority").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/authority/menu/role").
		SetDescription("apiDesc.menuAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST"),
	)

	// API

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/api/create").
		SetDescription("apiDesc.createApi").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/api/update").
		SetDescription("apiDesc.updateApi").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/api/delete").
		SetDescription("apiDesc.deleteAPI").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/api/list").
		SetDescription("apiDesc.APIList").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/api").
		SetDescription("apiDesc.getApiById").
		SetAPIGroup("api").
		SetMethod("POST"),
	)

	// DICTIONARY

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary/create").
		SetDescription("apiDesc.createDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary/update").
		SetDescription("apiDesc.updateDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary/delete").
		SetDescription("apiDesc.deleteDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary_detail/delete").
		SetDescription("apiDesc.deleteDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary_detail/create").
		SetDescription("apiDesc.createDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary_detail/update").
		SetDescription("apiDesc.updateDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary_detail/list").
		SetDescription("apiDesc.getDictionaryListDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dictionary/list").
		SetDescription("apiDesc.getDictionaryList").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict").
		SetDescription("apiDesc.getUserById").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	// OAUTH

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth_provider/create").
		SetDescription("apiDesc.createProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth_provider/update").
		SetDescription("apiDesc.updateProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth_provider/delete").
		SetDescription("apiDesc.deleteProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth_provider/list").
		SetDescription("apiDesc.geProviderList").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth/login").
		SetDescription("apiDesc.oauthLogin").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth_provider").
		SetDescription("apiDesc.getProviderById").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	// TOKEN

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token/create").
		SetDescription("apiDesc.createToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token/update").
		SetDescription("apiDesc.updateToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token/delete").
		SetDescription("apiDesc.deleteToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token/list").
		SetDescription("apiDesc.getTokenList").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token/logout").
		SetDescription("apiDesc.forceLoggingOut").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token").
		SetDescription("apiDesc.getTokenById").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	// DEPARTMENT

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department/create").
		SetDescription("apiDesc.createDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department/update").
		SetDescription("apiDesc.updateDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department/delete").
		SetDescription("apiDesc.deleteDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department/list").
		SetDescription("apiDesc.getDepartmentList").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department").
		SetDescription("apiDesc.getDepartmentById").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	// POSITION

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/position/create").
		SetDescription("apiDesc.createPosition").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/position/update").
		SetDescription("apiDesc.updatePosition").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/position/delete").
		SetDescription("apiDesc.deletePosition").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/position/list").
		SetDescription("apiDesc.getPositionList").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/position").
		SetDescription("apiDesc.getPositionById").
		SetAPIGroup("position").
		SetMethod("POST"),
	)

	// TASK
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task/create").
		SetDescription("apiDesc.createTask").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task/update").
		SetDescription("apiDesc.updateTask").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task/delete").
		SetDescription("apiDesc.deleteTask").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task/list").
		SetDescription("apiDesc.getTaskList").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task").
		SetDescription("apiDesc.getTaskById").
		SetAPIGroup("task").
		SetMethod("POST"),
	)

	// TASK LOG
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task_log/create").
		SetDescription("apiDesc.createTaskLog").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task_log/update").
		SetDescription("apiDesc.updateTaskLog").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task_log/delete").
		SetDescription("apiDesc.deleteTaskLog").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/task_log/list").
		SetDescription("apiDesc.getTaskLogList").
		SetAPIGroup("task_log").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
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
