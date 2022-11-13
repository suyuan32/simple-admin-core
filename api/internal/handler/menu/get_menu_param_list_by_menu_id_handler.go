package menu

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/menu"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /menu/param/list menu GetMenuParamListByMenuId
//
// Get menu extra parameters by menu ID | 获取某个菜单的额外参数列表
//
// Get menu extra parameters by menu ID | 获取某个菜单的额外参数列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: MenuParamListByMenuIdResp

func GetMenuParamListByMenuIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewGetMenuParamListByMenuIdLogic(r, svcCtx)
		resp, err := l.GetMenuParamListByMenuId(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
