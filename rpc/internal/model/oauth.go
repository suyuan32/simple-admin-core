package model

import (
	"gorm.io/gorm"
)

type OauthProvider struct {
	gorm.Model
	Name         string `json:"name,omitempty" gorm:"comment:the provider's name'"`
	ClientID     string `json:"clientID,omitempty"  gorm:"comment:the client id"`
	ClientSecret string `json:"clientSecret,omitempty" gorm:"comment:the client secret"`
	RedirectURL  string `json:"redirectURL,omitempty" gorm:"the redirect url"`
	Scopes       string `json:"scopes,omitempty" gorm:"comment:the scopes"`
	AuthURL      string `json:"authURL" gorm:"the auth url of the provider"`
	TokenURL     string `json:"tokenURL" gorm:"comment:he token url of the provider"`
	AuthStyle    int    `json:"authStyle" gorm:"comment:the auth style"`
}
