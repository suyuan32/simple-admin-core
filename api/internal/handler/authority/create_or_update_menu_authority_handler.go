package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/authority"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /authority/menu/create_or_update authority CreateOrUpdateMenuAuthority
//
// Create or update menu authorization information | 创建或更新菜单权限
//
// Create or update menu authorization information | 创建或更新菜单权限
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MenuAuthorityInfoReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateMenuAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuAuthorityInfoReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authority.NewCreateOrUpdateMenuAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrUpdateMenuAuthority(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
