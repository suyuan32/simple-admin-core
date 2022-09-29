package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name          string `json:"name" gorm:"comment:role name"`                                                        // role name | 角色名
	Value         string `json:"value" gorm:"unique;not null;comment: role value for permission control in front end"` // role value for permission control in front end | 角色值，用于前端权限控制
	DefaultRouter string `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"`                          // default menu : dashboard | 默认登录页面
	Status        uint32 `json:"status" gorm:"default:0;comment:role status:0 normal, 1 ban"`                          // status | 状态
	Remark        string `json:"remark" gorm:"comment:remark"`                                                         // remark | 备注
	OrderNo       uint32 `json:"orderNo" gorm:"default:0;comment:number for sorting"`                                  // order number | 排序编号
	Menus         []Menu `json:"menus" gorm:"many2many:role_menus"`
}
