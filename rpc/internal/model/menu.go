package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuLevel uint32      `json:"level" gorm:"comment:menu level"`                         // menu level | 菜单层级
	MenuType  uint32      `json:"type" gorm:"comment:menu type: 0. group 1. menu"`         // menu type | 菜单类型 （菜单或目录）0 目录 1 菜单
	ParentId  uint        `json:"parentId" gorm:"comment:parent menu id"`                  // parent menu ID | 父菜单ID
	Path      string      `json:"path" gorm:"comment:index path"`                          // index path | 菜单路由路径
	Name      string      `json:"name" gorm:"comment:index name"`                          // index name | 菜单名称
	Redirect  string      `json:"redirect" gorm:"comment: redirect path"`                  // redirect path | 跳转路径 （外链）
	Component string      `json:"component" gorm:"comment:the path of vue file"`           // the path of vue file | 组件路径
	OrderNo   uint32      `json:"orderNo" gorm:"default:0;comment:numbers for sorting"`    // sorting numbers | 排序编号
	Disabled  bool        `json:"disabled" gorm:"default:false;comment:if true, disable;"` // disable status | 是否停用
	Meta      Meta        `json:"meta" gorm:"embedded;comment:extra parameters"`           // extra parameters | 额外参数
	Roles     []Role      `json:"roles" gorm:"many2many:role_menus;"`
	Children  []Menu      `json:"children" gorm:"foreignKey:ParentId;references:id"`
	Param     []MenuParam `json:"parameters"`
}

type Meta struct {
	Title              string `json:"title" gorm:"comment:menu name"`                                          // menu name | 菜单显示标题
	Icon               string `json:"icon" gorm:"comment:menu icon"`                                           // menu icon | 菜单图标
	HideMenu           bool   `json:"hideMenu" gorm:"default:false;comment:hide the menu"`                     // hide menu | 是否隐藏菜单
	HideBreadcrumb     bool   `json:"hideBreadcrumb" gorm:"default:true;comment: hide the breadcrumb"`         // hide the breadcrumb | 隐藏面包屑
	CurrentActiveMenu  string `json:"currentActiveMenu" gorm:"comment:set the active menu"`                    // set the active menu | 激活菜单
	IgnoreKeepAlive    bool   `json:"ignoreKeepAlive" gorm:"comment: do not keep alive the tab"`               // do not keep alive the tab | 取消页面缓存
	HideTab            bool   `json:"hideTab" gorm:"comment: hide the tab header"`                             // hide the tab header | 隐藏页头
	FrameSrc           string `json:"frameSrc" gorm:"comment:iframe path"`                                     // show iframe | 内嵌 iframe
	CarryParam         bool   `json:"carryParam" gorm:"comment:the route carries parameters or not"`           // the route carries parameters or not | 携带参数
	HideChildrenInMenu bool   `json:"hideChildrenInMenu" gorm:"comment:hide children menu or not"`             // hide children menu or not | 隐藏所有子菜单
	Affix              bool   `json:"affix" gorm:"comment: affix tab"`                                         // affix tab | Tab 固定
	DynamicLevel       uint32 `json:"dynamicLevel" gorm:"the maximum number of pages the router can open"`     // the maximum number of pages the router can open | 能打开的子TAB数
	RealPath           string `json:"realPath" gorm:"comment:the real path of the route without dynamic part"` // the real path of the route without dynamic part | 菜单路由不包含参数部分
}

type MenuParam struct {
	gorm.Model
	MenuId uint
	Type   string `json:"type" gorm:"comment:pass parameters via params or query "` // pass parameters via params or query | 参数类型
	Key    string `json:"key" gorm:"comment:the key of parameters"`                 // the key of parameters | 参数键
	Value  string `json:"value" gorm:"comment:the value of parameters"`             // the value of parameters | 参数值
}
