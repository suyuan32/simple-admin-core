package middleware

import (
	"bytes"
	"github.com/duke-git/lancet/slice"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/pkg/enum"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type AuthorityMiddleware struct {
	Cbn   *casbin.Enforcer
	Rds   *redis.Redis
	Trans *i18n.Translator
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds *redis.Redis, trans *i18n.Translator) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn:   cbn,
		Rds:   rds,
		Trans: trans,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the path
		obj := r.URL.Path
		// get the method
		act := r.Method
		// get the role id
		roleIds := r.Context().Value("roleId").(string)

		// check jwt blacklist
		jwtResult, err := m.Rds.Get("token_" + r.Header.Get("Authorization"))
		if err != nil {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}
		if jwtResult == "1" {
			logx.Errorw("token in blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		result := batchCheck(m.Cbn, roleIds, act, obj, w, m.Rds)

		if result {
			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", r.Context().Value("userId").(string)),
				logx.Field("path", obj), logx.Field("method", act))
			// create intercepter，target：fill the "id" in r.Body
			createIntercepter(obj, r)
			next(w, r)
			return
		} else {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", roleIds),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, errorx.NewCodeError(enum.PermissionDenied, m.Trans.Trans(r.Header.Get("Accept-Language"),
				"common.permissionDeny")))
			return
		}
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds, act, obj string, w http.ResponseWriter, rds *redis.Redis) bool {
	var checkReq [][]any
	for _, v := range strings.Split(roleIds, ",") {
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

// check create operator
// waring：this rule must limit .api route,do not include 'create' without insert!
func createIntercepter(path string, r *http.Request) *http.Request {
	if slice.LastIndexOf(strings.Split(path, "/"), "create") == 1 {
		buffer, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(buffer))
		defer r.Body.Close()
		formValues := url.Values{}
		// fill the rpc must in field "id"
		formValues.Set("id", "0")
		formDataStr := formValues.Encode()
		formDataBytes := []byte(formDataStr)
		buffer = append(buffer, formDataBytes...)
		r.ContentLength = int64(len(buffer))
		r.Body = io.NopCloser(bytes.NewBuffer(buffer))
	}
	return r
}
