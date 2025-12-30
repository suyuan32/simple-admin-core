package warehouse

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/warehouse"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /warehouse/update warehouse UpdateWarehouse
//
// Update warehouse information | 更新仓库
//
// Update warehouse information | 更新仓库
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WarehouseInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateWarehouseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WarehouseInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := warehouse.NewUpdateWarehouseLogic(r.Context(), svcCtx)
		resp, err := l.UpdateWarehouse(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
