package i18n

import (
	"embed"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"

	"github.com/suyuan32/simple-admin-core/pkg/utils/errcode"
)

//go:embed locale/*.json
var localeFS embed.FS

type Translator struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func (l *Translator) NewBundle() {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFileFS(localeFS, "locale/zh.json")
	logx.Must(err)
	_, err = bundle.LoadMessageFileFS(localeFS, "locale/en.json")
	logx.Must(err)

	l.bundle = bundle
}

func (l *Translator) Trans(lang string, msgId string) string {
	localizer := i18n.NewLocalizer(l.bundle, lang)
	message, err := localizer.LocalizeMessage(&i18n.Message{ID: msgId})
	if err != nil {
		return msgId
	}
	return message
}

func (l *Translator) TransError(lang string, err error) error {
	localizer := i18n.NewLocalizer(l.bundle, lang)
	message, e := localizer.LocalizeMessage(&i18n.Message{ID: strings.Split(err.Error(), "desc = ")[1]})
	if e != nil || message == "" {
		message = err.Error()
	}

	if errcode.IsGrpcError(err) {
		return errorx.NewApiError(errcode.CodeFromGrpcError(err), message)
	} else if apiErr, ok := err.(*errorx.ApiError); ok {
		return errorx.NewApiError(apiErr.Code, message)
	} else {
		return errorx.NewApiError(http.StatusInternalServerError, message)
	}
}
