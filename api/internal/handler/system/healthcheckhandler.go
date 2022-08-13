package core

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/core"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewHealthCheckLogic(r.Context(), svcCtx)
		err := l.HealthCheck()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
