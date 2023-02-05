package position

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/position"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /position/create_or_update position CreateOrUpdatePosition
//
// Create or update position information | 创建或更新职位
//
// Create or update position information | 创建或更新职位
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdatePositionReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdatePositionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdatePositionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := position.NewCreateOrUpdatePositionLogic(r, svcCtx)
		resp, err := l.CreateOrUpdatePosition(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
