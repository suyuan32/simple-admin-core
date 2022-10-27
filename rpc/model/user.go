package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        string `json:"uuid" gorm:"comment:user's UUID"`                                     // UUID | 用户 UUID
	Username    string `json:"username" gorm:"unique;comment:user's login name'"`                   // user's login name | 登录名
	Password    string `json:"password"  gorm:"comment:password"`                                   // password | 密码
	Nickname    string `json:"nickName" gorm:"unique;comment:nickname"`                             // nickname | 昵称
	SideMode    string `json:"sideMode" gorm:"default:dark;comment:template mode"`                  // template mode | 布局方式
	Avatar      string `json:"Avatar" gorm:"comment:avatar"`                                        // avatar | 头像路径
	BaseColor   string `json:"baseColor" gorm:"default:#fff;comment:base color of template"`        // base color of template | 后台页面色调
	ActiveColor string `json:"activeColor" gorm:"default:#1890ff;comment:active color of template"` // active color of template | 当前激活的颜色设定
	RoleId      uint32 `json:"roleId" gorm:"default:2;comment:user's role id for access control"`   // role id | 角色ID
	Mobile      string `json:"mobile"  gorm:"comment:mobile number"`                                // mobile number | 手机号
	Email       string `json:"email"  gorm:"comment:email address"`                                 // email | 邮箱号
	Status      int32  `json:"status" gorm:"default:1;comment:user status 1: normal 2: ban"`        // status | 用户状态 1 正常 2 禁用
}
