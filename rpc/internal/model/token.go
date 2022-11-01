package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UUID      string    `json:"UUID" gorm:"index,comment:user's UUID"`              // User's UUID | 用户的UUID
	Token     string    `json:"Token" gorm:"comment: token string"`                 // Token string | Token 字符串
	Source    string    `json:"Source" gorm:"comment:log in source such as github"` // Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）
	Status    uint32    `json:"Status" gorm:"comment: JWT status 0 ban 1 active"`   // JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	ExpiredAt time.Time `json:"ExpiredAt" gorm:"index,comment:expire time"`         // Expire time | 过期时间
}
