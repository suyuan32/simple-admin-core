package model

import "gorm.io/gorm"

// Dictionary table struct
// 字典表结构
type Dictionary struct {
	gorm.Model
	Title  string             `json:"title,omitempty" yaml:"title" gorm:"comment:the title shown in the UI"`                                // the title shown in the ui | 展示名称 （建议配合i18n）
	Name   string             `json:"name,omitempty" yaml:"name" gorm:"index,unique,comment:the name of dictionary for search"`             // the name of dictionary for search | 字典搜索名称
	Status bool               `json:"status,omitempty" yaml:"status" gorm:"comment:the status of dictionary (true enable | false disable)"` // the status of dictionary (true enable | false disable) | 字典状态
	Desc   string             `json:"desc,omitempty" yaml:"desc" gorm:"comment:the descriptions of dictionary"`                             // the descriptions of dictionary | 字典描述
	Detail []DictionaryDetail `json:"detail" yaml:"detail"`
}

// DictionaryDetail table struct
// 字典键值表结构
type DictionaryDetail struct {
	gorm.Model
	Title        string `json:"title,omitempty" yaml:"title" gorm:"comment:the title shown in the UI"` // the title shown in the UI | 展示名
	Key          string `json:"key,omitempty" yaml:"key" gorm:"comment:key"`                           // key | 键
	Value        string `json:"value,omitempty" yaml:"value" gorm:"comment:value"`                     // value | 值
	Status       bool   `json:"status,omitempty" yaml:"status" gorm:"comment:status"`                  // status | 状态
	DictionaryID uint
}
