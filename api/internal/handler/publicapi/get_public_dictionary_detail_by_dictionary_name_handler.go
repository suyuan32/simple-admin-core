package publicapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicapi"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route get /dict/public/{name} publicapi GetPublicDictionaryDetailByDictionaryName
//
// Get dictionary detail by dictionary name without logging in | 无需登录通过字典名称获取字典内容
//
// Get dictionary detail by dictionary name without logging in | 无需登录通过字典名称获取字典内容
//
// Responses:
//  200: DictionaryDetailListResp

func GetPublicDictionaryDetailByDictionaryNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryNameReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicapi.NewGetPublicDictionaryDetailByDictionaryNameLogic(r.Context(), svcCtx)
		resp, err := l.GetPublicDictionaryDetailByDictionaryName(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
