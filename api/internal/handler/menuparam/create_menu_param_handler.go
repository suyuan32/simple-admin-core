package menuparam

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/menuparam"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /menu_param/create menuparam CreateMenuParam
//
// Create menu param information | 创建菜单参数
//
// Create menu param information | 创建菜单参数
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MenuParamInfo
//
// Responses:
//  200: BaseMsgResp

func CreateMenuParamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuParamInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menuparam.NewCreateMenuParamLogic(r.Context(), svcCtx)
		resp, err := l.CreateMenuParam(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
