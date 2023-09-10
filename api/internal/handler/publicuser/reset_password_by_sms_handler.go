package publicuser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicuser"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /user/reset_password_by_sms publicuser ResetPasswordBySms
//
// Reset password by Sms | 通过短信重置密码
//
// Reset password by Sms | 通过短信重置密码
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ResetPasswordBySmsReq
//
// Responses:
//  200: BaseMsgResp

func ResetPasswordBySmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPasswordBySmsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicuser.NewResetPasswordBySmsLogic(r.Context(), svcCtx)
		resp, err := l.ResetPasswordBySms(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
