package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /user/status user UpdateUserStatus
//
// Set user's status | 更新用户状态
//
// Set user's status | 更新用户状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StatusCodeUUIDReq
//
// Responses:
//  200: BaseMsgResp

func UpdateUserStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusCodeUUIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserStatusLogic(r, svcCtx)
		resp, err := l.UpdateUserStatus(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
