package department

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/department"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /department/create_or_update department CreateOrUpdateDepartment
//
// Create or update Department information | 创建或更新Department
//
// Create or update Department information | 创建或更新Department
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateDepartmentReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateDepartmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateDepartmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := department.NewCreateOrUpdateDepartmentLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateDepartment(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
