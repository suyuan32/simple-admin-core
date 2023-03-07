package oauthprovider

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/oauthprovider"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /oauth_provider/list oauthprovider GetOauthProviderList
//
// Get oauth provider list | 获取第三方列表
//
// Get oauth provider list | 获取第三方列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OauthProviderListReq
//
// Responses:
//  200: OauthProviderListResp

func GetOauthProviderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthProviderListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauthprovider.NewGetOauthProviderListLogic(r.Context(), svcCtx)
		resp, err := l.GetOauthProviderList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
