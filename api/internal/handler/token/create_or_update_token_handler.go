package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /token/create_or_update token CreateOrUpdateToken
//
// Create or update Token information | 创建或更新Token
//
// Create or update Token information | 创建或更新Token
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateTokenReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := token.NewCreateOrUpdateTokenLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateToken(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
