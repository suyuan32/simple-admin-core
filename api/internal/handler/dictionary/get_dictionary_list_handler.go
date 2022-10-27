package dictionary

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /dict/list dictionary GetDictionaryList
//
// Get dictionary list | 获取字典列表
//
// Get dictionary list | 获取字典列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryListReq
//
// Responses:
//  200: DictionaryListResp
//  401: SimpleMsg
//  500: SimpleMsg

func GetDictionaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewGetDictionaryListLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
