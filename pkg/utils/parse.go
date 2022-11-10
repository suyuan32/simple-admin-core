package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

func ParseTags(lang string) []language.Tag {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		logx.Error("parse language failed")
		return []language.Tag{language.Chinese}
	}

	return tags
}
