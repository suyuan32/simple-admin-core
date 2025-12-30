package stock_movement

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/stock_movement"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /stock_movement/list stock_movement GetStockMovementList
//
// Get stock movement list | 获取库存移动列表
//
// Get stock movement list | 获取库存移动列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StockMovementListReq
//
// Responses:
//  200: StockMovementListResp

func GetStockMovementListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockMovementListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := stock_movement.NewGetStockMovementListLogic(r.Context(), svcCtx)
		resp, err := l.GetStockMovementList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
