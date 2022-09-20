package svc

import (
	"testing"

	ja_lang "github.com/go-playground/locales/ja"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestAddTrans(t *testing.T) {
	jaLang := ja_lang.New()
	err := httpx.XValidator.Uni.AddTranslator(jaLang, true)
	if err != nil {
		t.Error(err.Error())
		return
	}
	jaTrans, _ := httpx.XValidator.Uni.GetTranslator("ja")
	httpx.XValidator.Trans["ja"] = jaTrans
	err = ja_translations.RegisterDefaultTranslations(httpx.XValidator.Validator, jaTrans)
	if err != nil {
		t.Error(err.Error())
		return
	}

	type User struct {
		Username string `validate:"alphanum,max=20"`
		Password string `validate:"min=6,max=30"`
	}
	u := User{
		Username: "admin",
		Password: "1",
	}
	result := httpx.XValidator.Validate(u, "ja")

	if result != "Passwordの長さは少なくとも6文字はなければなりません " {
		t.Error(result)
	}
}
