package token

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /token/list token GetTokenList
//
// Get Token list | 获取token列表
//
// Get Token list | 获取token列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TokenListReq
//
// Responses:
//  200: TokenListResp
//  401: SimpleMsg
//  500: SimpleMsg

func GetTokenListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TokenListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := token.NewGetTokenListLogic(r.Context(), svcCtx)
		resp, err := l.GetTokenList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
