package dictionary

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /dict/detail/list dictionary GetDetailByDictionaryName
//
// Get dictionary detail list by dictionary name | 根据字典名获取字典键值列表
//
// Get dictionary detail list by dictionary name | 根据字典名获取字典键值列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryDetailReq
//
// Responses:
//  200: DictionaryDetailListResp

func GetDetailByDictionaryNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewGetDetailByDictionaryNameLogic(r, svcCtx)
		resp, err := l.GetDetailByDictionaryName(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
