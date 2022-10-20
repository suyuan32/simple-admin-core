package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UUID     string    `json:"UUID" gorm:"index,comment:user's UUID"`              // user's UUID | 用户的UUID
	Token    string    `json:"Token" gorm:"comment: token string"`                 // Token string | Token 字符串
	Source   string    `json:"Source" gorm:"comment:log in source such as github"` // log in source such as github | Token 来源 （本地为core, 第三方如github等）
	Status   bool      `json:"Status" gorm:"comment: JWT status 0 ban 1 active"`   // JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	ExpireAt time.Time `json:"ExpireAt" gorm:"index,comment:expire time"`          // Expire time | 过期时间
}
