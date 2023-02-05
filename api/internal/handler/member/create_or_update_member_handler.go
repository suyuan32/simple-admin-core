package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/member"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /member/create_or_update member CreateOrUpdateMember
//
// Create or update member information | 创建或更新会员
//
// Create or update member information | 创建或更新会员
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMemberReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateMemberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member.NewCreateOrUpdateMemberLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateMember(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
