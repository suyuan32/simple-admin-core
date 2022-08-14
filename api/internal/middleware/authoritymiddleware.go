package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/common/errorx"

	"github.com/casbin/casbin/v2"
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
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		} else if roleStatus == "0" {
			httpx.Error(w, errorx.NewApiError(http.StatusBadRequest, errorx.RoleForbidden))
			return
		}

		sub := roleId
		result, err := m.Cbn.Enforce(sub, obj, act)
		if err != nil {
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, errorx.ApiRequestFailed))
			return
		}
		if result {
			next(w, r)
			return
		} else {
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		// use for temporary cancel authorization
		//next(w, r)
	}
}
