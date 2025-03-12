package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/rolectx"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-common/utils/jwt"
)

type AuthorityMiddleware struct {
	Cbn *casbin.Enforcer
	Rds redis.UniversalClient
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient) *AuthorityMiddleware {
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
		roleIds, err := rolectx.GetRoleIDFromCtx(r.Context())
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// check jwt blacklist
		jwtResult, err := m.Rds.Get(context.Background(), config.RedisTokenPrefix+jwt.StripBearerPrefixFromToken(r.Header.Get("Authorization"))).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}
		if jwtResult == "1" {
			logx.Errorw("token in blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		result := batchCheck(m.Cbn, roleIds, act, obj)

		if result {
			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", r.Context().Value("userId").(string)),
				logx.Field("path", obj), logx.Field("method", act))
			next(w, r)
			return
		} else {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", roleIds),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, errorx.NewApiForbiddenError("You do not have permission to access the API"))
			return
		}
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds []string, act, obj string) bool {
	var checkReq [][]any
	for _, v := range roleIds {
		checkReq = append(checkReq, []any{v, obj, act})
	}

	result, err := cbn.BatchEnforce(checkReq)
	if err != nil {
		logx.Errorw("Casbin enforce error", logx.Field("detail", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
