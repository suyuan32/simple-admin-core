package dictionary

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /dict dictionary CreateOrUpdateDictionary
//
// Create or update dictionary information | 创建或更新字典信息
//
// Create or update dictionary information | 创建或更新字典信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateDictionaryReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateDictionaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewCreateOrUpdateDictionaryLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateDictionary(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
