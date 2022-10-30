package model

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Path        string `json:"path" gorm:"index:idx_path,priority:1;comment:API path"`                    // API path | API 路径
	Description string `json:"description" gorm:"comment:API description"`                                // API description | API 描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:API group"`                                         // API group | API 分组
	Method      string `json:"method" gorm:"default:POST;index:idx_path,priority:2;comment: HTTP method"` // HTTP method | HTTP 请求类型
}
