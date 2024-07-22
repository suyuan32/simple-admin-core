package configuration

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/configuration"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /configuration/list configuration GetConfigurationList
//
// Get configuration list | 获取Configuration列表
//
// Get configuration list | 获取Configuration列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ConfigurationListReq
//
// Responses:
//  200: ConfigurationListResp

func GetConfigurationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConfigurationListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := configuration.NewGetConfigurationListLogic(r.Context(), svcCtx)
		resp, err := l.GetConfigurationList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
