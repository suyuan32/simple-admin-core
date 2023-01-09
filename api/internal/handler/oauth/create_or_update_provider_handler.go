package oauth

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/oauth"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /oauth/provider/create_or_update oauth CreateOrUpdateProvider
//
// Create or update Provider information | 创建或更新提供商
//
// Create or update Provider information | 创建或更新提供商
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateProviderReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateProviderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateProviderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauth.NewCreateOrUpdateProviderLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateProvider(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
