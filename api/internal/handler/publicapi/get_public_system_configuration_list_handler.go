package publicapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicapi"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
)

func GetPublicSystemConfigurationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := publicapi.NewGetPublicSystemConfigurationListLogic(r.Context(), svcCtx)
		resp, err := l.GetPublicSystemConfigurationList()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
