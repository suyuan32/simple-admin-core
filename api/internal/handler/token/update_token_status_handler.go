package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /token/status token UpdateTokenStatus
//
// Set token status | 设置token状态
//
// Set token status | 设置token状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StatusCodeUUIDReq
//
// Responses:
//  200: BaseMsgResp

func UpdateTokenStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusCodeUUIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := token.NewUpdateTokenStatusLogic(r, svcCtx)
		resp, err := l.UpdateTokenStatus(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
