package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /dictionary/update dictionary UpdateDictionary
//
// Update dictionary information | 更新字典
//
// Update dictionary information | 更新字典
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateDictionaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewUpdateDictionaryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDictionary(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
