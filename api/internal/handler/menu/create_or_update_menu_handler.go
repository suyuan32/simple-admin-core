package menu

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/menu"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /menu/create_or_update menu CreateOrUpdateMenu
//
// Create or update menu information | 创建或更新菜单
//
// Create or update menu information | 创建或更新菜单
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMenuReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewCreateOrUpdateMenuLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateMenu(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
