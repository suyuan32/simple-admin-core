package logic

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/util"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/casbin/casbin/v2"
	model2 "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

//  init database

func (l *InitDatabaseLogic) InitDatabase(in *core.Empty) (*core.BaseResp, error) {
	// add lock to avoid duplicate initialization
	lock := redis.NewRedisLock(l.svcCtx.Redis, "init_database_lock")
	lock.SetExpire(60)
	if ok, err := lock.Acquire(); !ok || err != nil {
		if !ok {
			logx.Error("last initialization is running")
			return nil, status.Error(codes.InvalidArgument, errorx.InitRunning)
		} else {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, status.Error(codes.Internal, errorx.RedisError)
		}
	}
	defer func() {
		recover()
		lock.Release()
	}()

	// judge if the initialization had been done
	var apis []model.Api
	check := l.svcCtx.DB.Find(&apis)
	if check.RowsAffected != 0 {
		err := l.svcCtx.Redis.Set("database_init_state", "1")
		if err != nil {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, status.Error(codes.Internal, errorx.RedisError)
		}
		return &core.BaseResp{Msg: errorx.AlreadyInit}, nil
	}

	// set default state value
	l.svcCtx.Redis.Setex("database_error_msg", "", 300)
	l.svcCtx.Redis.Set("database_init_state", "0")

	// initialize table structure
	err := l.svcCtx.DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Api{},
		&model.Menu{},
		&model.MenuParam{},
		&model.Dictionary{},
		&model.DictionaryDetail{},
		&model.OauthProvider{},
		&model.Token{},
	)

	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = l.insertUserData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertRoleData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertMenuData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertApiData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertRoleMenuAuthorityData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertCasbinPoliciesData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = l.insertProviderData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}

	l.svcCtx.Redis.Set("database_init_state", "1")
	return &core.BaseResp{Msg: errorx.Success}, nil
}

// insert init user data
func (l *InitDatabaseLogic) insertUserData() error {
	users := []model.User{
		{
			UUID:     uuid.NewString(),
			Username: "admin",
			Password: util.BcryptEncrypt("simple-admin"),
			Nickname: "Admin",
			Email:    "simple_admin@gmail.com",
			RoleId:   1,
		},
	}
	result := l.svcCtx.DB.CreateInBatches(users, 100)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	} else {
		return nil
	}
}

// insert init apis data
func (l *InitDatabaseLogic) insertRoleData() error {
	roles := []model.Role{
		{
			Name:          "sys.role.admin",
			Value:         "admin",
			DefaultRouter: "dashboard",
			Status:        1,
			Remark:        "超级管理员",
			OrderNo:       1,
		},
		{
			Name:          "sys.role.stuff",
			Value:         "stuff",
			DefaultRouter: "dashboard",
			Status:        1,
			Remark:        "普通员工",
			OrderNo:       2,
		},
		{
			Name:          "sys.role.member",
			Value:         "member",
			DefaultRouter: "dashboard",
			Status:        1,
			Remark:        "注册会员",
			OrderNo:       3,
		},
	}
	result := l.svcCtx.DB.CreateInBatches(roles, 100)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	} else {
		return nil
	}
}

// insert init user data
func (l *InitDatabaseLogic) insertApiData() error {
	apis := []model.Api{
		// user
		{
			Path:        "/user/login",
			Description: "apiDesc.userLogin",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/register",
			Description: "apiDesc.userRegister",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user",
			Description: "apiDesc.createOrUpdateUser",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/change-password",
			Description: "apiDesc.userChangePassword",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/info",
			Description: "apiDesc.userInfo",
			ApiGroup:    "user",
			Method:      "GET",
		},
		{
			Path:        "/user/list",
			Description: "apiDesc.userList",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user",
			Description: "apiDesc.deleteUser",
			ApiGroup:    "user",
			Method:      "DELETE",
		},
		{
			Path:        "/user/perm",
			Description: "apiDesc.userPermissions",
			ApiGroup:    "user",
			Method:      "GET",
		},
		{
			Path:        "/user/profile",
			Description: "apiDesc.userProfile",
			ApiGroup:    "user",
			Method:      "GET",
		},
		{
			Path:        "/user/profile",
			Description: "apiDesc.updateProfile",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/logout",
			Description: "apiDesc.logout",
			ApiGroup:    "user",
			Method:      "GET",
		},
		// role
		{
			Path:        "/role",
			Description: "apiDesc.createOrUpdateRole",
			ApiGroup:    "role",
			Method:      "POST",
		},
		{
			Path:        "/role",
			Description: "apiDesc.deleteRole",
			ApiGroup:    "role",
			Method:      "DELETE",
		},
		{
			Path:        "/role/list",
			Description: "apiDesc.roleList",
			ApiGroup:    "role",
			Method:      "POST",
		},
		{
			Path:        "/role/status",
			Description: "apiDesc.setRoleStatus",
			ApiGroup:    "role",
			Method:      "POST",
		},
		// menu
		{
			Path:        "/menu",
			Description: "apiDesc.createOrUpdateMenu",
			ApiGroup:    "menu",
			Method:      "POST",
		},
		{
			Path:        "/menu",
			Description: "apiDesc.deleteMenu",
			ApiGroup:    "menu",
			Method:      "DELETE",
		},
		{
			Path:        "/menu/list",
			Description: "apiDesc.menuList",
			ApiGroup:    "menu",
			Method:      "GET",
		},
		{
			Path:        "/menu/role",
			Description: "apiDesc.roleMenu",
			ApiGroup:    "menu",
			Method:      "GET",
		},
		{
			Path:        "/menu/param",
			Description: "apiDesc.createOrUpdateMenuParam",
			ApiGroup:    "menu",
			Method:      "POST",
		},
		{
			Path:        "/menu/param/list",
			Description: "apiDesc.menuParamListByMenuId",
			ApiGroup:    "menu",
			Method:      "POST",
		},
		{
			Path:        "/menu/param",
			Description: "apiDesc.deleteMenuParam",
			ApiGroup:    "menu",
			Method:      "DELETE",
		},
		// captcha
		{
			Path:        "/captcha",
			Description: "apiDesc.captcha",
			ApiGroup:    "captcha",
			Method:      "GET",
		},
		// authorization
		{
			Path:        "/authority/api",
			Description: "apiDesc.createOrUpdateApiAuthority",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		{
			Path:        "/authority/api/role",
			Description: "apiDesc.APIAuthorityOfRole",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		{
			Path:        "/authority/menu",
			Description: "apiDesc.createOrUpdateMenuAuthority",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		{
			Path:        "/authority/menu/role",
			Description: "apiDesc.menuAuthorityOfRole",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		// api
		{
			Path:        "/api",
			Description: "apiDesc.createOrUpdateApi",
			ApiGroup:    "api",
			Method:      "POST",
		},
		{
			Path:        "/api",
			Description: "apiDesc.deleteAPI",
			ApiGroup:    "api",
			Method:      "DELETE",
		},
		{
			Path:        "/api/list",
			Description: "apiDesc.APIList",
			ApiGroup:    "api",
			Method:      "POST",
		},
		// dictionary
		{
			Path:        "/dict",
			Description: "apiDesc.createOrUpdateDictionary",
			ApiGroup:    "dictionary",
			Method:      "POST",
		},
		{
			Path:        "/dict",
			Description: "apiDesc.deleteDictionary",
			ApiGroup:    "dictionary",
			Method:      "DELETE",
		},
		{
			Path:        "/dict/detail",
			Description: "apiDesc.deleteDictionaryDetail",
			ApiGroup:    "dictionary",
			Method:      "DELETE",
		},
		{
			Path:        "/dict/detail",
			Description: "apiDesc.createOrUpdateDictionaryDetail",
			ApiGroup:    "dictionary",
			Method:      "POST",
		},
		{
			Path:        "/dict/detail/list",
			Description: "apiDesc.getDictionaryListDetail",
			ApiGroup:    "dictionary",
			Method:      "POST",
		},
		{
			Path:        "/dict/list",
			Description: "apiDesc.getDictionaryList",
			ApiGroup:    "dictionary",
			Method:      "POST",
		},
		// oauth APIs
		{
			Path:        "/oauth/provider",
			Description: "apiDesc.createOrUpdateProvider",
			ApiGroup:    "oauth",
			Method:      "POST",
		},
		{
			Path:        "/oauth/provider",
			Description: "apiDesc.deleteProvider",
			ApiGroup:    "oauth",
			Method:      "DELETE",
		},
		{
			Path:        "/oauth/provider/list",
			Description: "apiDesc.geProviderList",
			ApiGroup:    "oauth",
			Method:      "POST",
		},
		{
			Path:        "/oauth/login",
			Description: "apiDesc.oauthLogin",
			ApiGroup:    "oauth",
			Method:      "POST",
		},
		// token api
		{
			Path:        "/token",
			Description: "apiDesc.createOrUpdateToken",
			ApiGroup:    "token",
			Method:      "POST",
		},
		{
			Path:        "/token",
			Description: "apiDesc.deleteToken",
			ApiGroup:    "token",
			Method:      "DELETE",
		},
		{
			Path:        "/token/list",
			Description: "apiDesc.getTokenList",
			ApiGroup:    "token",
			Method:      "POST",
		},
		{
			Path:        "/token/status",
			Description: "apiDesc.setTokenStatus",
			ApiGroup:    "token",
			Method:      "POST",
		},
		{
			Path:        "/token/logout",
			Description: "sys.user.forceLoggingOut",
			ApiGroup:    "token",
			Method:      "POST",
		},
	}
	result := l.svcCtx.DB.CreateInBatches(apis, 100)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	} else {
		return nil
	}
}

// init menu data
func (l *InitDatabaseLogic) insertMenuData() error {
	menus := []model.Menu{
		{
			Model:     gorm.Model{ID: 1},
			MenuLevel: 0,
			MenuType:  0,
			ParentId:  1,
			Path:      "",
			Name:      "root",
			Component: "",
			OrderNo:   0,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "",
				Icon:               "",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  0,
			ParentId:  1,
			Path:      "/dashboard",
			Name:      "Dashboard",
			Component: "/dashboard/workbench/index",
			OrderNo:   0,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.dashboard.dashboard",
				Icon:               "ant-design:home-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				CurrentActiveMenu:  "",
				IgnoreKeepAlive:    false,
				HideTab:            false,
				FrameSrc:           "",
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  0,
			ParentId:  1,
			Path:      "",
			Name:      "System Management",
			Component: "LAYOUT",
			OrderNo:   1,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.systemManagementTitle",
				Icon:               "ant-design:tool-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/menu",
			Name:      "MenuManagement",
			Component: "/sys/menu/index",
			OrderNo:   1,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.menuManagementTitle",
				Icon:               "ant-design:bars-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/role",
			Name:      "Role Management",
			Component: "/sys/role/index",
			OrderNo:   2,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.roleManagementTitle",
				Icon:               "ant-design:user-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/api",
			Name:      "API Management",
			Component: "/sys/api/index",
			OrderNo:   4,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.apiManagementTitle",
				Icon:               "ant-design:api-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/user",
			Name:      "User Management",
			Component: "/sys/user/index",
			OrderNo:   3,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.userManagementTitle",
				Icon:               "ant-design:user-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  1,
			ParentId:  1,
			Path:      "/file",
			Name:      "File Management",
			Component: "/file/index",
			OrderNo:   2,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.fileManagementTitle",
				Icon:               "ant-design:folder-open-outlined",
				HideMenu:           true,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/dictionary",
			Name:      "Dictionary Management",
			Component: "/sys/dictionary/index",
			OrderNo:   5,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.dictionaryManagementTitle",
				Icon:               "ant-design:book-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  0,
			ParentId:  1,
			Path:      "",
			Name:      "Other Pages",
			Component: "LAYOUT",
			OrderNo:   4,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.otherPages",
				Icon:               "ant-design:question-circle-outlined",
				HideMenu:           true,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  10,
			Path:      "/dictionary/detail",
			Name:      "Dictionary Detail",
			Component: "/sys/dictionary/detail",
			OrderNo:   1,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.dictionaryDetailManagementTitle",
				Icon:               "ant-design:align-left-outlined",
				HideMenu:           true,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  1,
			ParentId:  10,
			Path:      "/profile",
			Name:      "Profile",
			Component: "/sys/profile/index",
			OrderNo:   3,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.userProfileTitle",
				Icon:               "ant-design:profile-outlined",
				HideMenu:           true,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/oauth",
			Name:      "Oauth Management",
			Component: "/sys/oauth/index",
			OrderNo:   6,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.oauthManagement",
				Icon:               "ant-design:unlock-filled",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/token",
			Name:      "Token Management",
			Component: "/sys/token/index",
			OrderNo:   7,
			Disabled:  false,
			Meta: model.Meta{
				Title:              "routes.system.tokenManagement",
				Icon:               "ant-design:lock-outlined",
				HideMenu:           false,
				HideBreadcrumb:     true,
				IgnoreKeepAlive:    false,
				HideTab:            false,
				CarryParam:         false,
				HideChildrenInMenu: false,
				Affix:              false,
				DynamicLevel:       20,
			},
		},
	}
	result := l.svcCtx.DB.CreateInBatches(menus, 100)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	} else {
		return nil
	}
}

// insert admin menu authority

func (l *InitDatabaseLogic) insertRoleMenuAuthorityData() error {
	var menus []model.Menu
	result := l.svcCtx.DB.Find(&menus)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	}

	var insertString strings.Builder
	insertString.WriteString("insert into role_menus values ")
	for i := 0; i < len(menus); i++ {
		if i != len(menus)-1 {
			insertString.WriteString(fmt.Sprintf("(%d, %d),", menus[i].ID, 1))
		} else {
			insertString.WriteString(fmt.Sprintf("(%d, %d);", menus[i].ID, 1))
		}
	}

	result = l.svcCtx.DB.Exec(insertString.String())
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	} else {
		return nil
	}
}

// init casbin policies

func (l *InitDatabaseLogic) insertCasbinPoliciesData() error {
	var apis []model.Api
	result := l.svcCtx.DB.Find(&apis)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	}

	var policies [][]string
	for _, v := range apis {
		policies = append(policies, []string{"1", v.Path, v.Method})
	}

	csb := getCasbin(l.svcCtx.DB)
	addResult, err := csb.AddPolicies(policies)

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if addResult {
		return nil
	} else {
		return status.Error(codes.Internal, err.Error())
	}
}

func getCasbin(db *gorm.DB) *casbin.SyncedEnforcer {
	var syncedEnforcer *casbin.SyncedEnforcer
	a, _ := gormadapter.NewAdapterByDB(db)
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model2.NewModelFromString(text)
	if err != nil {
		log.Fatal("InitCasbin: import model fail!", err)
		return nil
	}
	syncedEnforcer, err = casbin.NewSyncedEnforcer(m, a)
	if err != nil {
		log.Fatal("InitCasbin: NewSyncedEnforcer fail!", err)
		return nil
	}
	err = syncedEnforcer.LoadPolicy()
	if err != nil {
		log.Fatal("InitCasbin: LoadPolicy fail!", err)
		return nil
	}
	return syncedEnforcer
}

func (l *InitDatabaseLogic) insertProviderData() error {
	providers := []model.OauthProvider{
		{
			Name:         "google",
			ClientID:     "your client id",
			ClientSecret: "your client secret",
			RedirectURL:  "http://localhost:3100/oauth/login/callback",
			Scopes:       "email openid",
			AuthURL:      "https://accounts.google.com/o/oauth2/auth",
			TokenURL:     "https://oauth2.googleapis.com/token",
			AuthStyle:    1,
			InfoURL:      "https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
		},
		{
			Name:         "github",
			ClientID:     "your client id",
			ClientSecret: "your client secret",
			RedirectURL:  "http://localhost:3100/oauth/login/callback",
			Scopes:       "email openid",
			AuthURL:      "https://github.com/login/oauth/authorize",
			TokenURL:     "https://github.com/login/oauth/access_token",
			AuthStyle:    2,
			InfoURL:      "https://api.github.com/user",
		},
	}
	result := l.svcCtx.DB.CreateInBatches(providers, 10)
	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return status.Error(codes.Internal, result.Error.Error())
	} else {
		return nil
	}
}
