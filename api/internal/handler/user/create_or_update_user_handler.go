package user

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /user user CreateOrUpdateUser
//
// Create or update user's information | 新增或更新用户
//
// Create or update user's information | 新增或更新用户
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateUserReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewCreateOrUpdateUserLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateUser(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
