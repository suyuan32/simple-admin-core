package inventory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/inventory"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /inventory/list inventory GetInventoryList
//
// Get inventory list | 获取库存列表
//
// Get inventory list | 获取库存列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: InventoryListReq
//
// Responses:
//  200: InventoryListResp

func GetInventoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InventoryListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := inventory.NewGetInventoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetInventoryList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
