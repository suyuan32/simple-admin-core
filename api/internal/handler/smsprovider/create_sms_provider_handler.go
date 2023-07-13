package smsprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/smsprovider"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sms_provider/create smsprovider CreateSmsProvider
//
// Create sms provider information | 创建短信配置
//
// Create sms provider information | 创建短信配置
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SmsProviderInfo
//
// Responses:
//  200: BaseMsgResp

func CreateSmsProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := smsprovider.NewCreateSmsProviderLogic(r.Context(), svcCtx)
		resp, err := l.CreateSmsProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
