package dictionarydetail

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /dictionary_detail/update dictionarydetail UpdateDictionaryDetail
//
// Update dictionary detail information | 更新字典键值
//
// Update dictionary detail information | 更新字典键值
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DictionaryDetailInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryDetailInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionarydetail.NewUpdateDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDictionaryDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
