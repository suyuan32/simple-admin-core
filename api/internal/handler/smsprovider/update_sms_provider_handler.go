package smsprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/smsprovider"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sms_provider/update smsprovider UpdateSmsProvider
//
// Update sms provider information | 更新短信配置
//
// Update sms provider information | 更新短信配置
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SmsProviderInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateSmsProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := smsprovider.NewUpdateSmsProviderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSmsProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
