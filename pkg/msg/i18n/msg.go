package i18n

const (
	// user service messages

	// UserAlreadyExists
	// EN: Username or email address had been registered
	// ZH_CN: 用户名或者邮箱已被注册
	UserAlreadyExists string = "sys.login.signupUserExist"

	// UserNotExists
	// EN: User is not registered
	// ZH_CN: 用户不存在
	UserNotExists string = "sys.login.userNotExist"

	// WrongCaptcha
	// EN: Wrong captcha
	// ZH_CN: 验证码错误
	WrongCaptcha string = "sys.login.wrongCaptcha"

	// WrongUsernameOrPassword
	// EN: Wrong username or password
	// ZH_CN: 用户名或密码错误
	WrongUsernameOrPassword string = "sys.login.wrongUsernameOrPassword"

	// menu service messages

	// ChildrenExistError
	// EN: Please delete menu's children first
	// ZH_CN: 请先删除子菜单
	ChildrenExistError string = "sys.menu.deleteChildrenDesc"

	// ParentNotExist
	// EN: The parent does not exist
	// ZH_CN: 父级不存在
	ParentNotExist string = "sys.menu.parentNotExist"

	// MenuNotExists
	// EN: Menu does not exist
	// ZH_CN: 菜单不存在
	MenuNotExists string = "sys.menu.menuNotExists"

	// MenuAlreadyExists
	// EN: Menu already exists
	// ZH_CN: 菜单已存在
	MenuAlreadyExists string = "sys.menu.menuAlreadyExists"

	// role service messages

	// DuplicateRoleValue
	// EN: Duplicate role value
	// ZH_CN: 角色值重复
	DuplicateRoleValue string = "sys.role.duplicateRoleValue"

	// UserExists
	// EN: Please delete users who belong to this role
	// ZH_CN: 请先删除该角色下的用户
	UserExists string = "sys.role.userExists"

	// RoleForbidden
	// EN: Your role is forbidden
	// ZH_CN: 您的角色已停用
	RoleForbidden string = "sys.role.roleForbidden"

	// WrongPassword
	// EN: Wrong Password
	// ZH_CN: 密码错误
	WrongPassword string = "sys.user.wrongPassword"

	// ApiAlreadyExists
	// En: API already exists
	// ZH_CN: API已存在
	ApiAlreadyExists string = "sys.apis.apiExists"

	// DictionaryAlreadyExists
	// En: Dictionary already exists
	// ZH_CN: 字典已存在
	DictionaryAlreadyExists string = "sys.dict.dictExists"

	// CreateAccountFirst
	// En: Please register an account with this email or bind the email to an account
	// ZH_CH: 请创建一个该邮箱的账号或绑定该邮箱到一个账号
	CreateAccountFirst string = "sys.oauth.createAccount"
)
