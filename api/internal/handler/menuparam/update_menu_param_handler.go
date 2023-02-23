package menuparam

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/menuparam"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /menu_param/update menuparam UpdateMenuParam
//
// Update menu param information | 更新菜单参数
//
// Update menu param information | 更新菜单参数
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MenuParamInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateMenuParamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuParamInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menuparam.NewUpdateMenuParamLogic(r, svcCtx)
		resp, err := l.UpdateMenuParam(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
