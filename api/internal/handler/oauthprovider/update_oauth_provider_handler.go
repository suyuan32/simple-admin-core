package oauthprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/oauthprovider"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /oauth_provider/update oauthprovider UpdateOauthProvider
//
// Update oauth provider information | 更新第三方
//
// Update oauth provider information | 更新第三方
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OauthProviderInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateOauthProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthProviderInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauthprovider.NewUpdateOauthProviderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateOauthProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
