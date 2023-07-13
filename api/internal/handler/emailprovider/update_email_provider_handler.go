package emailprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/emailprovider"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /email_provider/update emailprovider UpdateEmailProvider
//
// Update email provider information | 更新邮箱服务配置
//
// Update email provider information | 更新邮箱服务配置
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmailProviderInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateEmailProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emailprovider.NewUpdateEmailProviderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateEmailProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
