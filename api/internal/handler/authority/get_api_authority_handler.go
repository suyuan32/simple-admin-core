package authority

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/authority"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /authority/api/role authority GetApiAuthority
//
// Get role's API authorization list | 获取角色api权限列表
//
// Get role's API authorization list | 获取角色api权限列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: ApiAuthorityListResp
//  401: SimpleMsg
//  500: SimpleMsg

func GetApiAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authority.NewGetApiAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.GetApiAuthority(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
