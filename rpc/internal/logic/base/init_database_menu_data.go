package base

import (
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
)

// insert initial menu data
func (l *InitDatabaseLogic) insertMenuData() error {
	var menus []*ent.MenuCreate

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(common.DefaultParentId).
		SetPath("/dashboard_dir").
		SetName("DashBoardDir").
		SetComponent("LAYOUT").
		SetSort(1).
		SetTitle("route.managementCenter").
		SetIcon("ic:baseline-view-in-ar").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(1).
		SetPath("/dashboard").
		SetName("Dashboard").
		SetComponent("/dashboard/workbench/index").
		SetSort(0).
		SetTitle("route.dashboard").
		SetIcon("ant-design:home-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(common.DefaultParentId).
		SetPath("/system").
		SetName("SystemManagement").
		SetComponent("LAYOUT").
		SetSort(999).
		SetTitle("route.systemManagementTitle").
		SetIcon("ant-design:tool-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/menu").
		SetName("MenuManagement").
		SetComponent("/sys/menu/index").
		SetSort(1).
		SetTitle("route.menuManagementTitle").
		SetIcon("ant-design:bars-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/role").
		SetName("RoleManagement").
		SetComponent("/sys/role/index").
		SetSort(2).
		SetTitle("route.roleManagementTitle").
		SetIcon("ant-design:user-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/user").
		SetName("UserManagement").
		SetComponent("/sys/user/index").
		SetSort(3).
		SetTitle("route.userManagementTitle").
		SetIcon("ant-design:user-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/department").
		SetName("DepartmentManagement").
		SetComponent("/sys/department/index").
		SetSort(4).
		SetTitle("route.departmentManagement").
		SetIcon("ic:outline-people-alt").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/api").
		SetName("APIManagement").
		SetComponent("/sys/api/index").
		SetSort(5).
		SetTitle("route.apiManagementTitle").
		SetIcon("ant-design:api-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/dictionary").
		SetName("DictionaryManagement").
		SetComponent("/sys/dictionary/index").
		SetSort(6).
		SetHideChildrenInMenu(true).
		SetTitle("route.dictionaryManagementTitle").
		SetIcon("ant-design:book-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(common.DefaultParentId).
		SetPath("/other").
		SetName("OtherPages").
		SetComponent("LAYOUT").
		SetSort(1000).
		SetTitle("route.otherPages").
		SetIcon("ant-design:question-circle-outlined").
		SetHideMenu(true).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(10).
		SetPath("/profile").
		SetName("Profile").
		SetComponent("/sys/profile/index").
		SetSort(3).
		SetTitle("route.userProfileTitle").
		SetIcon("ant-design:profile-outlined").
		SetHideMenu(true).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/oauth").
		SetName("OauthManagement").
		SetComponent("/sys/oauth/index").
		SetSort(6).
		SetTitle("route.oauthManagement").
		SetIcon("ant-design:unlock-filled").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/token").
		SetName("TokenManagement").
		SetComponent("/sys/token/index").
		SetSort(7).
		SetTitle("route.tokenManagement").
		SetIcon("ant-design:lock-outlined").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/position").
		SetName("PositionManagement").
		SetComponent("/sys/position/index").
		SetSort(8).
		SetTitle("route.positionManagement").
		SetIcon("ic:twotone-work-outline").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/task").
		SetName("TaskManagement").
		SetComponent("/sys/task/index").
		SetSort(8).
		SetTitle("route.taskManagement").
		SetIcon("ic:baseline-access-alarm").
		SetHideMenu(true).
		SetServiceName("Job"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/configuration").
		SetName("ConfigurationManagement").
		SetComponent("/sys/configuration/index").
		SetSort(9).
		SetTitle("route.configurationManagement").
		SetIcon("carbon:data-2").
		SetHideMenu(false).
		SetServiceName("Core"),
	)

	menus = append(menus, l.svcCtx.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(common.DefaultParentId).
		SetPath("/module_store").
		SetName("ModuleStore").
		SetComponent("IFrame").
		SetSort(99).
		SetTitle("route.moduleStore").
		SetIcon("ion:bag-handle-outline").
		SetHideMenu(false).
		SetServiceName("Core").
		SetFrameSrc("https://simple-admin.simple-hub.top/#/store/index"),
	)

	err := l.svcCtx.DB.Menu.CreateBulk(menus...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}
