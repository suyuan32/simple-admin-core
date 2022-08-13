package model

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Path        string `json:"path" gorm:"comment:api path"`                    // api path
	Description string `json:"description" gorm:"comment:api description"`      // api description
	ApiGroup    string `json:"apiGroup" gorm:"comment:api group"`               // api group
	Method      string `json:"method" gorm:"default:POST;comment: http method"` // http method
}
