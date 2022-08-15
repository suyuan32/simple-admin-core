package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuLevel uint32      `json:"level" gorm:"comment:menu level"`
	MenuType  uint32      `json:"type" gorm:"comment:menu type: 1. group 2. menu"`
	ParentId  uint        `json:"parentId" gorm:"comment:parent menu id"`                  // parent menu id
	Path      string      `json:"path" gorm:"comment:index path"`                          // index path
	Name      string      `json:"name" gorm:"comment:index name"`                          // index name
	Redirect  string      `json:"redirect" gorm:"comment: redirect path"`                  // redirect path
	Component string      `json:"component" gorm:"comment:the path of vue file"`           // the path of vue file
	OrderNo   uint32      `json:"orderNo" gorm:"default:0;comment:numbers for sorting"`    // sorting numbers
	Disabled  bool        `json:"disabled" gorm:"default:false;comment:if true, disable;"` //disable status
	Meta      Meta        `json:"meta" gorm:"embedded;comment:extra parameters"`           // extra parameters
	Roles     []Role      `json:"roles" gorm:"many2many:role_menus;"`
	Children  []Menu      `json:"children" gorm:"foreignKey:ParentId;references:id"`
	Param     []MenuParam `json:"parameters"`
}

type Meta struct {
	KeepAlive         bool   `json:"keepAlive" gorm:"default:false;comment:save in cache?"`           // save in cache
	HideMenu          bool   `json:"hideMenu" gorm:"default:false;comment:hide the menu"`             // hide menu
	HideBreadcrumb    bool   `json:"hideBreadcrumb" gorm:"default:true;comment: hide the breadcrumb"` // hide the breadcrumb
	CurrentActiveMenu string `json:"currentActiveMenu" gorm:"comment:set the active menu"`            // set the active menu
	Title             string `json:"title" gorm:"comment:menu name"`                                  // menu name
	Icon              string `json:"icon" gorm:"comment:menu icon"`                                   // menu icon
	CloseTab          bool   `json:"closeTab" gorm:"default:true;comment:auto close tab"`             // auto close tab
}

type MenuParam struct {
	gorm.Model
	MenuId uint
	Type   string `json:"type" gorm:"comment:pass parameters via params or query "` // pass parameters via params or query
	Key    string `json:"key" gorm:"comment:the key of parameters"`                 // the key of parameters
	Value  string `json:"value" gorm:"comment:the value of parameters"`             // the value of parameters
}
