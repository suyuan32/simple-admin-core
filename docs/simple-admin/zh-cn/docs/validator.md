## Validator 使用

> 只需要在 api 中结构声明中使用 validate tag 即可实现校验

```text
 // login request | 登录参数
    // swagger:model LoginReq
    LoginReq {
        // User Name | 用户名
        // Required: true
        // Max length: 20
        Username   string `json:"username" validate:"alphanum,max=20"`

        // Password | 密码
        // Required: true
        // Min length: 6
        // Max length: 30
        Password   string `json:"password" validate:"max=30,min=6"`

        // Captcha Id which store in redis | 验证码编号, 存在redis中
        // Required: true
        // Max length: 20
        CaptchaId  string `json:"captchaId"  validate:"len=20"`

        // The Captcha which users input | 用户输入的验证码
        // Required: true
        // Max length: 5
        Captcha    string `json:"captcha" validate:"len=5"`
    }
```

> 支持多语言，默认支持中文和英文，如果需要其他语言请自行添加

[例子](https://github.com/suyuan32/simple-admin-core/blob/master/api/internal/svc/servicecontext_test.go)

```go
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

```