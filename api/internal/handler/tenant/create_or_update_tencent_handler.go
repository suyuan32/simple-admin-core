package tenant

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/tenant"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /tencent/create_or_update tenant CreateOrUpdateTencent
//

//

//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateTenantReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateTencentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateTenantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tenant.NewCreateOrUpdateTencentLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateTencent(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
