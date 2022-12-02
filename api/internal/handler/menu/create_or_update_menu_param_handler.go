package menu

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/menu"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /menu/param/create_or_update menu CreateOrUpdateMenuParam
//
// Create or update menu parameters | 创建或更新菜单参数
//
// Create or update menu parameters | 创建或更新菜单参数
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMenuParamReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateMenuParamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMenuParamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewCreateOrUpdateMenuParamLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateMenuParam(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
