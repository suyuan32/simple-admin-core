package stock_movement

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/stock_movement"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /stock_movement/update stock_movement UpdateStockMovement
//
// Update stock movement information | 更新库存移动
//
// Update stock movement information | 更新库存移动
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StockMovementInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateStockMovementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockMovementInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := stock_movement.NewUpdateStockMovementLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStockMovement(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
