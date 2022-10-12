# 权限验证

我们使用Casbin进行权限验证.

> 初始化

```go
cbn := utils.NewCasbin(db)
```

[cbn](https://github.com/suyuan32/simple-admin-core/blob/master/api/internal/svc/servicecontext.go)

> 创建中间件

```go
package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/suyuan32/simple-admin-core/common/logmessage"
	"github.com/suyuan32/simple-admin-core/common/message"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthorityMiddleware struct {
	Cbn *casbin.SyncedEnforcer
	Rds *redis.Redis
}

func NewAuthorityMiddleware(cbn *casbin.SyncedEnforcer, rds *redis.Redis) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn: cbn,
		Rds: rds,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the path
		obj := r.URL.Path
		// get the method
		act := r.Method
		// get the role id
		roleId := r.Context().Value("roleId").(json.Number).String()
		// check the role status
		roleStatus, err := m.Rds.Hget("roleData", fmt.Sprintf("%s_status", roleId))
		if err != nil {
			logx.Errorw(logmessage.RedisError, logx.Field("Detail", err.Error()))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		} else if roleStatus == "0" {
			logx.Errorw("Role is on forbidden status", logx.Field("RoleId", roleId))
			httpx.Error(w, errorx.NewApiError(http.StatusBadRequest, message.RoleForbidden))
			return
		}

		sub := roleId
		result, err := m.Cbn.Enforce(sub, obj, act)
		if err != nil {
			logx.Errorw("Casbin enforce error", logx.Field("Detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, errorx.ApiRequestFailed))
			return
		}
		if result {
			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", r.Context().Value("userId").(string)),
				logx.Field("Path", obj), logx.Field("Method", act))
			next(w, r)
			return
		} else {
			logx.Errorw("The role is not permitted to access the API", logx.Field("RoleId", roleId),
				logx.Field("Path", obj), logx.Field("Method", act))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}
	}
}

```

[Middleware](https://github.com/suyuan32/simple-admin-core/blob/master/api/internal/middleware/authoritymiddleware.go)

> 在 api 文件中使用

```text
@server(
    jwt: Auth
    group: api
    middleware: Authority
)
```

[如何使用中间件](https://go-zero.dev/cn/docs/advance/middleware)