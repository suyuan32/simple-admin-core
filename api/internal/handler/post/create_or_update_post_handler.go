package post

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/post"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /post/create_or_update post CreateOrUpdatePost
//
// Create or update post information | 创建或更新职位
//
// Create or update post information | 创建或更新职位
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdatePostReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdatePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdatePostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := post.NewCreateOrUpdatePostLogic(r, svcCtx)
		resp, err := l.CreateOrUpdatePost(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
