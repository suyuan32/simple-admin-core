package warehouse

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/warehouse"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /warehouse warehouse GetWarehouseById
//
// Get warehouse by ID | 通过ID获取仓库
//
// Get warehouse by ID | 通过ID获取仓库
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDReq
//
// Responses:
//  200: WarehouseInfoResp

func GetWarehouseByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := warehouse.NewGetWarehouseByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetWarehouseById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
