package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	l := &Translator{}
	l.NewBundle(LocaleFS)
	l.NewTranslator()
	res := l.Trans("zh", "login.userNotExist")
	assert.Equal(t, "用户不存在", res)
}
