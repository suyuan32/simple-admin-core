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
// Get oauth provider list | 获取OauthProvider列表
//
// Get oauth provider list | 获取OauthProvider列表
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
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauthprovider.NewGetOauthProviderListLogic(r, svcCtx)
		resp, err := l.GetOauthProviderList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
