package i18n

import (
	"fmt"
	"testing"
)

func TestTranslator(t *testing.T) {
	l := &Translator{}
	l.NewBundle()
	res := l.Trans("zh", "login.userNotExist")
	fmt.Println(res)
}
