package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name          string `json:"name" gorm:"comment:role name"`
	Value         string `json:"value" gorm:"unique;not null;comment: role value for permission control in front end"`
	DefaultRouter string `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` // default menu : dashboard
	Status        uint32 `json:"status" gorm:"default:0;comment:role status:0 normal, 1 ban"`
	Remark        string `json:"remark" gorm:"comment:remark"`
	OrderNo       uint32 `json:"orderNo" gorm:"default:0;comment:number for sorting"`
	Menus         []Menu `json:"menus" gorm:"many2many:role_menus"`
}
