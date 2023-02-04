package base

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *core.Empty) (*core.BaseResp, error) {
	// If your mysql speed is high, comment the code below.
	// Because the context deadline will reach if the database is too slow
	l.ctx = context.Background()

	// add lock to avoid duplicate initialization
	lock := redis.NewRedisLock(l.svcCtx.Redis, "init_database_lock")
	lock.SetExpire(60)
	if ok, err := lock.Acquire(); !ok || err != nil {
		if !ok {
			logx.Error("last initialization is running")
			return nil, statuserr.NewInternalError(i18n.InitRunning)
		} else {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.RedisError)
		}
	}
	defer func() {
		recover()
		lock.Release()
	}()

	// initialize table structure
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}

	// judge if the initialization had been done
	check, err := l.svcCtx.DB.API.Query().Count(l.ctx)

	if check != 0 {
		err := l.svcCtx.Redis.Set("database_init_state", "1")
		if err != nil {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.RedisError)
		}
		return &core.BaseResp{Msg: i18n.AlreadyInit}, nil
	}

	// set default state value
	l.svcCtx.Redis.Setex("database_error_msg", "", 300)
	l.svcCtx.Redis.Setex("database_init_state", "0", 300)

	err = l.insertUserData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}
	err = l.insertRoleData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}
	err = l.insertMenuData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}
	err = l.insertApiData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}
	err = l.insertRoleMenuAuthorityData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}
	err = l.insertCasbinPoliciesData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}

	err = l.insertProviderData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}

	err = l.insertDepartmentData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, statuserr.NewInternalError(err.Error())
	}

	l.svcCtx.Redis.Setex("database_init_state", "1", 300)
	return &core.BaseResp{Msg: i18n.Success}, nil
}

// insert init user data
func (l *InitDatabaseLogic) insertUserData() error {
	var users []*ent.UserCreate
	users = append(users, l.svcCtx.DB.User.Create().
		SetUsername("admin").
		SetNickname("admin").
		SetPassword(utils.BcryptEncrypt("simple-admin")).
		SetEmail("simple_admin@gmail.com").
		SetRoleID(1),
	)

	err := l.svcCtx.DB.User.CreateBulk(users...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert init apis data
func (l *InitDatabaseLogic) insertRoleData() error {
	var roles []*ent.RoleCreate
	roles = append(roles, l.svcCtx.DB.Role.Create().
		SetName("role.admin").
		SetValue("admin").
		SetRemark("超级管理员").
		SetSort(1),
	)

	roles = append(roles, l.svcCtx.DB.Role.Create().
		SetName("role.stuff").
		SetValue("stuff").
		SetRemark("普通员工").
		SetSort(2),
	)

	roles = append(roles, l.svcCtx.DB.Role.Create().
		SetName("role.member").
		SetValue("member").
		SetRemark("注册会员").
		SetSort(3),
	)

	err := l.svcCtx.DB.Role.CreateBulk(roles...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert init user data
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
		SetPath("/user/create_or_update").
		SetDescription("apiDesc.createOrUpdateUser").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/user/change-password").
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
		SetPath("/user/batch_delete").
		SetDescription("apiDesc.batchDeleteUser").
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
		SetPath("/user/status").
		SetDescription("apiDesc.updateUserStatus").
		SetAPIGroup("user").
		SetMethod("POST"),
	)

	// ROLE
	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/role/create_or_update").
		SetDescription("apiDesc.createOrUpdateRole").
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
		SetPath("/role/status").
		SetDescription("apiDesc.setRoleStatus").
		SetAPIGroup("role").
		SetMethod("POST"),
	)

	// MENU

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/create_or_update").
		SetDescription("apiDesc.createOrUpdateMenu").
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
		SetPath("/menu/role").
		SetDescription("apiDesc.roleMenu").
		SetAPIGroup("menu").
		SetMethod("GET"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/param/create_or_update").
		SetDescription("apiDesc.createOrUpdateMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/param/list").
		SetDescription("apiDesc.menuParamListByMenuId").
		SetAPIGroup("menu").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/menu/param/delete").
		SetDescription("apiDesc.deleteMenuParam").
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
		SetPath("/api/create_or_update").
		SetDescription("apiDesc.createOrUpdateApi").
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

	// DICTIONARY

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict/create_or_update").
		SetDescription("apiDesc.createOrUpdateDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict/delete").
		SetDescription("apiDesc.deleteDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict/detail/delete").
		SetDescription("apiDesc.deleteDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict/detail/create_or_update").
		SetDescription("apiDesc.createOrUpdateDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict/detail/list").
		SetDescription("apiDesc.getDictionaryListDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/dict/list").
		SetDescription("apiDesc.getDictionaryList").
		SetAPIGroup("dictionary").
		SetMethod("POST"),
	)

	// OAUTH

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth/provider/create_or_update").
		SetDescription("apiDesc.createOrUpdateProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth/provider/delete").
		SetDescription("apiDesc.deleteProvider").
		SetAPIGroup("oauth").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/oauth/provider/list").
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

	// TOKEN

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/token/create_or_update").
		SetDescription("apiDesc.createOrUpdateToken").
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
		SetPath("/token/status").
		SetDescription("apiDesc.setTokenStatus").
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
		SetPath("/token/batch_delete").
		SetDescription("apiDesc.batchDeleteToken").
		SetAPIGroup("token").
		SetMethod("POST"),
	)

	// department

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department/create_or_update").
		SetDescription("apiDesc.createOrUpdateDepartment").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	apis = append(apis, l.svcCtx.DB.API.Create().
		SetPath("/department/batch_delete").
		SetDescription("apiDesc.batchDeleteDepartment").
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
		SetPath("/department/status").
		SetDescription("apiDesc.updateDepartmentStatus").
		SetAPIGroup("department").
		SetMethod("POST"),
	)

	err := l.svcCtx.DB.API.CreateBulk(apis...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// init menu data
func (l *InitDatabaseLogic) insertMenuData() error {
	var menus []*ent.MenuCreate

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(0).
		SetPath("/dashboard").
		SetName("Dashboard").
		SetComponent("/dashboard/workbench/index").
		SetSort(0).
		SetTitle("route.dashboard").
		SetIcon("ant-design:home-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(0).
		SetPath("").
		SetName("System Management").
		SetComponent("LAYOUT").
		SetSort(1).
		SetTitle("route.systemManagementTitle").
		SetIcon("ant-design:tool-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/menu").
		SetName("MenuManagement").
		SetComponent("/sys/menu/index").
		SetSort(1).
		SetTitle("route.menuManagementTitle").
		SetIcon("ant-design:bars-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/role").
		SetName("Role Management").
		SetComponent("/sys/role/index").
		SetSort(2).
		SetTitle("route.roleManagementTitle").
		SetIcon("ant-design:user-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/api").
		SetName("API Management").
		SetComponent("/sys/api/index").
		SetSort(4).
		SetTitle("route.apiManagementTitle").
		SetIcon("ant-design:api-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/user").
		SetName("User Management").
		SetComponent("/sys/user/index").
		SetSort(3).
		SetTitle("route.userManagementTitle").
		SetIcon("ant-design:user-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(0).
		SetPath("/file").
		SetName("File Management").
		SetComponent("/file/index").
		SetSort(2).
		SetTitle("route.fileManagementTitle").
		SetIcon("ant-design:folder-open-outlined").
		SetHideMenu(true),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/dictionary").
		SetName("Dictionary Management").
		SetComponent("/sys/dictionary/index").
		SetSort(5).
		SetTitle("route.dictionaryManagementTitle").
		SetIcon("ant-design:book-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(0).
		SetPath("").
		SetName("Other Pages").
		SetComponent("LAYOUT").
		SetSort(4).
		SetTitle("route.otherPages").
		SetIcon("ant-design:question-circle-outlined").
		SetHideMenu(true),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(2).
		SetParentID(9).
		SetPath("/dictionary/detail").
		SetName("Dictionary Detail").
		SetComponent("/sys/dictionary/detail").
		SetSort(1).
		SetTitle("route.dictionaryDetailManagementTitle").
		SetIcon("ant-design:align-left-outlined").
		SetHideMenu(true),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(9).
		SetPath("/profile").
		SetName("Profile").
		SetComponent("/sys/profile/index").
		SetSort(3).
		SetTitle("route.userProfileTitle").
		SetIcon("ant-design:profile-outlined").
		SetHideMenu(true),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/oauth").
		SetName("Oauth Management").
		SetComponent("/sys/oauth/index").
		SetSort(6).
		SetTitle("route.oauthManagement").
		SetIcon("ant-design:unlock-filled").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/token").
		SetName("Token Management").
		SetComponent("/sys/token/index").
		SetSort(7).
		SetTitle("route.tokenManagement").
		SetIcon("ant-design:lock-outlined").
		SetHideMenu(false),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(2).
		SetPath("/department").
		SetName("Department Management").
		SetComponent("/sys/department/index").
		SetSort(8).
		SetTitle("route.departmentManagement").
		SetIcon("ant-design:lock-outlined").
		SetHideMenu(false),
	)

	err := l.svcCtx.DB.Menu.CreateBulk(menus...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert admin menu authority

func (l *InitDatabaseLogic) insertRoleMenuAuthorityData() error {
	count, err := l.svcCtx.DB.Menu.Query().Count(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	}

	var menuIds []uint64
	menuIds = make([]uint64, count)

	for i := range menuIds {
		menuIds[i] = uint64(i + 1)
	}

	err = l.svcCtx.DB.Role.Update().AddMenuIDs(menuIds...).Exec(l.ctx)

	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// init casbin policies

func (l *InitDatabaseLogic) insertCasbinPoliciesData() error {
	apis, err := l.svcCtx.DB.API.Query().All(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	}

	var policies [][]string
	for _, v := range apis {
		policies = append(policies, []string{"1", v.Path, v.Method})
	}

	csb, err := l.svcCtx.Config.CasbinConf.NewCasbin(l.svcCtx.Config.DatabaseConf.Type,
		l.svcCtx.Config.DatabaseConf.GetDSN())
	if err != nil {
		logx.Error("initialize casbin policy failed")
		return statuserr.NewInternalError(err.Error())
	}

	addResult, err := csb.AddPolicies(policies)
	if err != nil {
		return statuserr.NewInternalError(err.Error())
	}

	if addResult {
		return nil
	} else {
		return statuserr.NewInternalError(err.Error())
	}
}

func (l *InitDatabaseLogic) insertProviderData() error {
	var providers []*ent.OauthProviderCreate

	providers = append(providers, l.svcCtx.DB.OauthProvider.Create().
		SetName("google").
		SetClientID("your client id").
		SetClientSecret("your client secret").
		SetRedirectURL("http://localhost:3100/oauth/login/callback").
		SetScopes("email openid").
		SetAuthURL("https://accounts.google.com/o/oauth2/auth").
		SetTokenURL("https://oauth2.googleapis.com/token").
		SetAuthStyle(1).
		SetInfoURL("https://www.googleapis.com/oauth2/v2/userinfo?access_token="),
	)

	providers = append(providers, l.svcCtx.DB.OauthProvider.Create().
		SetName("github").
		SetClientID("your client id").
		SetClientSecret("your client secret").
		SetRedirectURL("http://localhost:3100/oauth/login/callback").
		SetScopes("email openid").
		SetAuthURL("https://github.com/login/oauth/authorize").
		SetTokenURL("https://github.com/login/oauth/access_token").
		SetAuthStyle(2).
		SetInfoURL("https://api.github.com/user"),
	)

	err := l.svcCtx.DB.OauthProvider.CreateBulk(providers...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert init department data
func (l *InitDatabaseLogic) insertDepartmentData() error {
	var departments []*ent.DepartmentCreate
	departments = append(departments, l.svcCtx.DB.Department.Create().
		SetName("department.managementDepartment").
		SetAncestors("").
		SetLeader("admin").
		SetEmail("simpleadmin@gmail.com").
		SetPhone("18888888888").
		SetRemark("Super Administrator").
		SetSort(1).
		SetParentID(0),
	)

	err := l.svcCtx.DB.Department.CreateBulk(departments...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return statuserr.NewInternalError(err.Error())
	} else {
		return nil
	}
}
