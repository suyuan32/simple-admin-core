package role

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/role"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /role/status role UpdateRoleStatus
//
// Set role status | 设置角色状态
//
// Set role status | 设置角色状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StatusCodeReq
//
// Responses:
//  200: SimpleMsg
//  401: SimpleMsg
//  500: SimpleMsg

func UpdateRoleStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewUpdateRoleStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRoleStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
