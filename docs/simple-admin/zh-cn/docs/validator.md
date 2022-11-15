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

[文件](https://github.com/suyuan32/simple-admin-tools/blob/master/rest/httpx/util.go)

```go
func NewValidator() *Validator {
	v := Validator{}
	en := enLang.New()
	zh := zhLang.New()
	v.Uni = ut.New(zh, en, zh)
	v.Validator = validator.New()
	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")
	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans
	// add support languages
	initSupportLanguages()

	err := enTranslations.RegisterDefaultTranslations(v.Validator, enTrans)
	if err != nil {
		logx.Errorw("register English translation failed", logx.Field("detail", err.Error()))
		return nil
	}
	err = zhTranslations.RegisterDefaultTranslations(v.Validator, zhTrans)
	if err != nil {
		logx.Errorw("register Chinese translation failed", logx.Field("detail", err.Error()))

		return nil
	}

	return &v
}

```

> 注意： 添加validate标签后默认不能为空，若需要允许为空需要添加 omitempty 

```go
// Get token list request params | token列表请求参数
    // swagger:model TokenListReq
    TokenListReq {
        PageInfo
        // User's UUID | 用户的UUID
        // Required: true
        // Max Length: 36
        UUID      string `json:"UUID" validate:"omitempty,len=36"`

        // user's nickname | 用户的昵称
        // Required: true
        // Max length: 10
        Nickname  string  `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`

        // User Name | 用户名
        // Required: true
        // Max length: 20
        Username   string `json:"username" validate:"omitempty,alphanum,max=20"`

        // The user's email address | 用户的邮箱
        // Required: true
        // Max length: 100
        Email     string `json:"email" validate:"omitempty,email,max=100"`
    }
```