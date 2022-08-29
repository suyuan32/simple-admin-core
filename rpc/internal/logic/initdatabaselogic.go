package logic

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
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
			logx.Error("Last initialization is running")
			return nil, status.Error(codes.InvalidArgument, errorx.InitRunning)
		} else {
			logx.Errorw(logmessage.RedisError, logx.Field("Detail", err.Error()))
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
			logx.Errorw(logmessage.RedisError, logx.Field("Detail", err.Error()))
			return nil, status.Error(codes.Internal, errorx.RedisError)
		}
		return &core.BaseResp{Msg: errorx.AlreadyInit}, nil
	}

	// set default state value
	l.svcCtx.Redis.Setex("database_error_msg", "", 300)
	l.svcCtx.Redis.Set("database_init_state", "0")

	// initialize table structure
	err := l.svcCtx.DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Api{}, &model.Menu{}, &model.MenuParam{})
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = l.insertUserData()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertRoleData()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertMenuData()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertApiData()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertRoleMenuAuthorityData()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
		l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = l.insertCasbinPoliciesData()
	if err != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", err.Error()))
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
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
			Description: "api_desc.userLogin",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/register",
			Description: "api_desc.userRegister",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/change-password",
			Description: "api_desc.userChangePassword",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/info",
			Description: "api_desc.userInfo",
			ApiGroup:    "user",
			Method:      "GET",
		},
		{
			Path:        "/user/list",
			Description: "api_desc.userList",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user",
			Description: "api_desc.deleteUser",
			ApiGroup:    "user",
			Method:      "DELETE",
		},
		{
			Path:        "/user/modify",
			Description: "api_desc.userModify",
			ApiGroup:    "user",
			Method:      "POST",
		},
		{
			Path:        "/user/perm",
			Description: "api_desc.userPermissions",
			ApiGroup:    "user",
			Method:      "GET",
		},
		{
			Path:        "/user/profile",
			Description: "api_desc.userProfile",
			ApiGroup:    "user",
			Method:      "GET",
		},
		{
			Path:        "/user/profile",
			Description: "api_desc.updateProfile",
			ApiGroup:    "user",
			Method:      "POST",
		},
		// role
		{
			Path:        "/role",
			Description: "api_desc.createOrUpdateRole",
			ApiGroup:    "role",
			Method:      "POST",
		},
		{
			Path:        "/role",
			Description: "api_desc.deleteRole",
			ApiGroup:    "role",
			Method:      "DELETE",
		},
		{
			Path:        "/role/list",
			Description: "api_desc.roleList",
			ApiGroup:    "role",
			Method:      "POST",
		},
		{
			Path:        "/role/status",
			Description: "api_desc.setRoleStatus",
			ApiGroup:    "role",
			Method:      "POST",
		},
		// menu
		{
			Path:        "/menu",
			Description: "api_desc.createOrUpdateMenu",
			ApiGroup:    "menu",
			Method:      "POST",
		},
		{
			Path:        "/menu",
			Description: "api_desc.deleteMenu",
			ApiGroup:    "menu",
			Method:      "DELETE",
		},
		{
			Path:        "/menu/list",
			Description: "api_desc.menuList",
			ApiGroup:    "menu",
			Method:      "GET",
		},
		{
			Path:        "/menu/role",
			Description: "api_desc.roleMenu",
			ApiGroup:    "menu",
			Method:      "GET",
		},
		// captcha
		{
			Path:        "/captcha",
			Description: "api_desc.captcha",
			ApiGroup:    "captcha",
			Method:      "GET",
		},
		// authorization
		{
			Path:        "/authority/api",
			Description: "api_desc.createOrUpdateApiAuthority",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		{
			Path:        "/authority/api/role",
			Description: "api_desc.APIAuthorityOfRole",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		{
			Path:        "/authority/menu",
			Description: "api_desc.createOrUpdateMenuAuthority",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		{
			Path:        "/authority/menu/role",
			Description: "api_desc.menuAuthorityOfRole",
			ApiGroup:    "authority",
			Method:      "POST",
		},
		// api
		{
			Path:        "/api",
			Description: "api_desc.createOrUpdateApi",
			ApiGroup:    "api",
			Method:      "POST",
		},
		{
			Path:        "/api",
			Description: "api_desc.deleteAPI",
			ApiGroup:    "api",
			Method:      "DELETE",
		},
		{
			Path:        "/api/list",
			Description: "api_desc.APIList",
			ApiGroup:    "api",
			Method:      "POST",
		},
	}
	result := l.svcCtx.DB.CreateInBatches(apis, 100)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
				KeepAlive:         false,
				HideMenu:          true,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "",
				Icon:              "",
				CloseTab:          false,
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
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.dashboard.dashboard",
				Icon:              "ant-design:home-outlined",
				CloseTab:          false,
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
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.systemManagementTitle",
				Icon:              "ant-design:tool-outlined",
				CloseTab:          false,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/menu_management",
			Name:      "MenuManagement",
			Component: "/sys/menu/index",
			OrderNo:   1,
			Disabled:  false,
			Meta: model.Meta{
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.menuManagementTitle",
				Icon:              "ant-design:bars-outlined",
				CloseTab:          false,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/role_management",
			Name:      "Role Management",
			Component: "/sys/role/index",
			OrderNo:   2,
			Disabled:  false,
			Meta: model.Meta{
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.roleManagementTitle",
				Icon:              "ant-design:user-outlined",
				CloseTab:          false,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/api_management",
			Name:      "API Management",
			Component: "/sys/api/index",
			OrderNo:   4,
			Disabled:  false,
			Meta: model.Meta{
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.apiManagementTitle",
				Icon:              "ant-design:api-outlined",
				CloseTab:          false,
			},
		},
		{
			MenuLevel: 2,
			MenuType:  1,
			ParentId:  3,
			Path:      "/user_management",
			Name:      "User Management",
			Component: "/sys/user/index",
			OrderNo:   3,
			Disabled:  false,
			Meta: model.Meta{
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.userManagementTitle",
				Icon:              "ant-design:user-outlined",
				CloseTab:          false,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  1,
			ParentId:  1,
			Path:      "/file_management",
			Name:      "File Management",
			Component: "/file/index",
			OrderNo:   2,
			Disabled:  false,
			Meta: model.Meta{
				KeepAlive:         false,
				HideMenu:          false,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.fileManagementTitle",
				Icon:              "ant-design:folder-open-outlined",
				CloseTab:          false,
			},
		},
		{
			MenuLevel: 1,
			MenuType:  1,
			ParentId:  1,
			Path:      "/profile",
			Name:      "Profile",
			Component: "/sys/profile/index",
			OrderNo:   3,
			Disabled:  false,
			Meta: model.Meta{
				KeepAlive:         false,
				HideMenu:          true,
				HideBreadcrumb:    true,
				CurrentActiveMenu: "",
				Title:             "routes.system.userProfileTitle",
				Icon:              "ant-design:profile-outlined",
				CloseTab:          false,
			},
		},
	}
	result := l.svcCtx.DB.CreateInBatches(menus, 100)
	if result.Error != nil {
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
		logx.Errorw(logmessage.DatabaseError, logx.Field("Detail", result.Error.Error()))
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
