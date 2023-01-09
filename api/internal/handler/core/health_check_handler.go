package core

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/core"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /core/health core HealthCheck
//
// Check the system status | 检查系统状态
//
// Check the system status | 检查系统状态
//

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewHealthCheckLogic(r, svcCtx)
		err := l.HealthCheck()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
