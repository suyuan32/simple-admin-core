package inventory

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/inventory"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /inventory/update inventory UpdateInventory
//
// Update inventory information | 更新库存
//
// Update inventory information | 更新库存
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: InventoryInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateInventoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InventoryInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := inventory.NewUpdateInventoryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInventory(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
