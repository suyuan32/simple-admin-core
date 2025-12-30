package stock_movement

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/stock_movement"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /stock_movement/create stock_movement CreateStockMovement
//
// Create stock movement information | 创建库存移动
//
// Create stock movement information | 创建库存移动
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StockMovementInfo
//
// Responses:
//  200: BaseMsgResp

func CreateStockMovementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockMovementInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := stock_movement.NewCreateStockMovementLogic(r.Context(), svcCtx)
		resp, err := l.CreateStockMovement(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
