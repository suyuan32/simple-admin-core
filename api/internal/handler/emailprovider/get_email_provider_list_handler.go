package emailprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/emailprovider"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /email_provider/list emailprovider GetEmailProviderList
//
// Get email provider list | 获取邮箱服务配置列表
//
// Get email provider list | 获取邮箱服务配置列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmailProviderListReq
//
// Responses:
//  200: EmailProviderListResp

func GetEmailProviderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailProviderListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emailprovider.NewGetEmailProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailProviderList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
