package base

import (
	"context"
	"errors"
	"github.com/bsm/redislock"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(_ *core.Empty) (*core.BaseResp, error) {
	// If your mysql speed is high, comment the code below.
	// Because the context deadline will reach if the database is too slow
	l.ctx = context.Background()

	// add lock to avoid duplicate initialization
	locker := redislock.New(l.svcCtx.Redis)

	lock, err := locker.Obtain(l.ctx, "INIT:DATABASE:LOCK", 10*time.Minute, nil)
	if errors.Is(err, redislock.ErrNotObtained) {
		logx.Error("last initialization is running")
		return nil, errorx.NewInternalError("i18n.InitRunning")
	} else if err != nil {
		logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
		return nil, errorx.NewInternalError("failed to get redis lock")
	}

	defer lock.Release(l.ctx)

	// initialize table structure
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", err.Error(), 300*time.Second).Err()
		return nil, errorx.NewInternalError(err.Error())
	}

	// force update casbin policy to avoid super administrator cannot log in when updating policy failed
	err = l.insertCasbinPoliciesData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", err.Error(), 300*time.Second).Err()
		return nil, errorx.NewInternalError(err.Error())
	}

	// judge if the initialization had been done
	check, err := l.svcCtx.DB.API.Query().Count(l.ctx)

	if check != 0 {
		err := l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:STATE", "1", 24*time.Hour).Err()
		if err != nil {
			logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
			return nil, errorx.NewInternalError(i18n.RedisError)
		}
		return &core.BaseResp{Msg: i18n.AlreadyInit}, nil
	}

	// set default state value
	_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", "", 300*time.Second)
	_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:STATE", "0", 300*time.Second)

	errHandler := func(err error) (*core.BaseResp, error) {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", err.Error(), 300*time.Second)
		return nil, errorx.NewInternalError(err.Error())
	}

	err = l.insertRoleData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertUserData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertMenuData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertApiData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertRoleMenuAuthorityData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertProviderData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertDepartmentData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertPositionData()
	if err != nil {
		return errHandler(err)
	}

	err = l.insertCasbinPoliciesData()
	if err != nil {
		return errHandler(err)
	}

	_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:STATE", "1", 24*time.Hour)
	return &core.BaseResp{Msg: i18n.Success}, nil
}

// insert init user data
func (l *InitDatabaseLogic) insertUserData() error {
	var users []*ent.UserCreate
	users = append(users, l.svcCtx.DB.User.Create().
		SetUsername("admin").
		SetNickname("admin").
		SetPassword(encrypt.BcryptEncrypt("simple-admin")).
		SetEmail("simple_admin@gmail.com").
		AddRoleIDs(1).
		SetDepartmentID(1).
		AddPositionIDs(1),
	)

	err := l.svcCtx.DB.User.CreateBulk(users...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert initial role data
func (l *InitDatabaseLogic) insertRoleData() error {
	var roles []*ent.RoleCreate
	roles = append(roles, l.svcCtx.DB.Role.Create().
		SetName("role.admin").
		SetCode("001").
		SetRemark("超级管理员").
		SetSort(1),
	)

	roles = append(roles, l.svcCtx.DB.Role.Create().
		SetName("role.stuff").
		SetCode("002").
		SetRemark("普通员工").
		SetSort(2),
	)

	err := l.svcCtx.DB.Role.CreateBulk(roles...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	}

	return nil
}

// insert initial admin menu authority data
func (l *InitDatabaseLogic) insertRoleMenuAuthorityData() error {
	count, err := l.svcCtx.DB.Menu.Query().Count(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	}

	var menuIds []uint64
	menuIds = make([]uint64, count)

	for i := range menuIds {
		menuIds[i] = uint64(i + 1)
	}

	err = l.svcCtx.DB.Role.Update().AddMenuIDs(menuIds...).Exec(l.ctx)

	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert initial Casbin policies
func (l *InitDatabaseLogic) insertCasbinPoliciesData() error {
	apis, err := l.svcCtx.DB.API.Query().All(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	}

	var roleCode string

	if adminData, roleErr := l.svcCtx.DB.Role.Query().Where(role.NameEQ("role.admin")).First(l.ctx); roleErr == nil {
		roleCode = adminData.Code
	} else {
		roleCode = "001"
	}

	var policies [][]string
	for _, v := range apis {
		policies = append(policies, []string{roleCode, v.Path, v.Method})
	}

	csb, err := l.svcCtx.Config.CasbinConf.NewCasbin(l.svcCtx.Config.DatabaseConf.Type,
		l.svcCtx.Config.DatabaseConf.GetDSN())

	if err != nil {
		logx.Error("initialize casbin policy failed")
		return errorx.NewInternalError(err.Error())
	}

	// clear old policy belongs to super admin
	var oldPolicies [][]string
	oldPolicies, err = csb.GetFilteredPolicy(0, roleCode)
	if err != nil {
		logx.Error("failed to get old Casbin policy", logx.Field("detail", err))
		return errorx.NewInternalError(err.Error())
	}

	if len(oldPolicies) != 0 {
		removeResult, err := csb.RemoveFilteredPolicy(0, roleCode)
		if err != nil {
			logx.Errorw("failed to remove roles policy", logx.Field("roleCode", roleCode), logx.Field("detail", err.Error()))
			return errorx.NewInternalError(err.Error())
		}
		if !removeResult {
			return errorx.NewInternalError("casbin.removeFailed")
		}
	}

	addResult, err := csb.AddPolicies(policies)
	if err != nil {
		return errorx.NewInternalError(err.Error())
	}

	if addResult {
		return nil
	} else {
		return errorx.NewInternalError(err.Error())
	}
}

// insert initial provider data
func (l *InitDatabaseLogic) insertProviderData() error {
	var providers []*ent.OauthProviderCreate

	providers = append(providers, l.svcCtx.DB.OauthProvider.Create().
		SetName("google").
		SetClientID("your client id").
		SetClientSecret("your client secret").
		SetRedirectURL("http://localhost:3100/oauth/login/callback").
		SetScopes("email openid").
		SetAuthURL("https://accounts.google.com/o/oauth2/auth").
		SetTokenURL("https://oauth2.googleapis.com/token").
		SetAuthStyle(1).
		SetInfoURL("https://www.googleapis.com/oauth2/v2/userinfo?access_token=TOKEN"),
	)

	providers = append(providers, l.svcCtx.DB.OauthProvider.Create().
		SetName("github").
		SetClientID("your client id").
		SetClientSecret("your client secret").
		SetRedirectURL("http://localhost:3100/oauth/login/callback").
		SetScopes("email openid").
		SetAuthURL("https://github.com/login/oauth/authorize").
		SetTokenURL("https://github.com/login/oauth/access_token").
		SetAuthStyle(2).
		SetInfoURL("https://api.github.com/user"),
	)

	err := l.svcCtx.DB.OauthProvider.CreateBulk(providers...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert initial department data
func (l *InitDatabaseLogic) insertDepartmentData() error {
	var departments []*ent.DepartmentCreate
	departments = append(departments, l.svcCtx.DB.Department.Create().
		SetName("department.managementDepartment").
		SetAncestors("").
		SetLeader("admin").
		SetEmail("simpleadmin@gmail.com").
		SetPhone("18888888888").
		SetRemark("Super Administrator").
		SetSort(1).
		SetParentID(common.DefaultParentId),
	)

	err := l.svcCtx.DB.Department.CreateBulk(departments...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}

// insert initial position data
func (l *InitDatabaseLogic) insertPositionData() error {
	var posts []*ent.PositionCreate
	posts = append(posts, l.svcCtx.DB.Position.Create().
		SetName("position.ceo").
		SetRemark("CEO").SetCode("001").SetSort(1),
	)

	err := l.svcCtx.DB.Position.CreateBulk(posts...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return errorx.NewInternalError(err.Error())
	} else {
		return nil
	}
}
