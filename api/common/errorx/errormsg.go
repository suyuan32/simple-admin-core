package errorx

// these messages are used for show error info in front end with different languages
// 多语言提示信息配置
const (
	// internal error

	// DatabaseError
	// normal database error
	DatabaseError string = "database error occur"

	// RedisError
	// normal redis error
	RedisError string = "redis error occur"

	// request error

	// ApiRequestFailed
	// EN: The interface request failed, please try again later!
	// ZH_CN: 请求出错，请稍候重试
	ApiRequestFailed string = "sys.api.apiRequestFailed"

	// CreateSuccess
	// EN: Create successfully
	// ZH_CN: 新建成功
	CreateSuccess string = "common.createSuccess"

	// CreateFailed
	// EN: Create failed
	// ZH_CN: 新建失败
	CreateFailed string = "common.createFailed"

	// UpdateSuccess
	// EN: Update successfully
	// ZH_CN: 更新成功
	UpdateSuccess string = "common.updateSuccess"

	// UpdateFailed
	// EN: Update failed
	// ZH_CN: 更新失败
	UpdateFailed string = "common.updateFailed"

	// DeleteSuccess
	// EN: Delete successfully
	// ZH_CN: 删除成功
	DeleteSuccess string = "common.deleteSuccess"

	// DeleteFailed
	// EN: Delete failed
	// ZH_CN: 删除失败
	DeleteFailed string = "common.deleteFailed"

	// GetInfoSuccess
	// EN: Get information Successfully
	// ZH_CN: 获取信息成功
	GetInfoSuccess string = "common.getInfoSuccess"

	// GetInfoFailed
	// EN: Get information Failed
	// ZH_CN: 获取信息失败
	GetInfoFailed string = "common.getInfoFailed"

	// TargetNotExist
	// EN: Target does not exist
	// ZH_CN: 目标不存在
	TargetNotExist string = "common.targetNotExist"

	// Success
	// EN: Successful
	// ZH_CN: 成功
	Success string = "common.successful"

	// Failed
	// EN: Failed
	// ZH_CN: 失败
	Failed string = "common.failed"

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

	// InitRunning
	// EN: The initialization is running...
	// ZH_CN: 正在初始化...
	InitRunning string = "sys.init.initializeIsRunning"

	// AlreadyInit
	// EN: The database had been initialized.
	// ZH_CN: 数据库已被初始化。
	AlreadyInit string = "sys.init.alreadyInit"
)
