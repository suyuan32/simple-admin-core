package member

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/member"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /member/status member UpdateMemberStatus
//
// Set member's status | 更新会员状态
//
// Set member's status | 更新会员状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StatusCodeUUIDReq
//
// Responses:
//  200: BaseMsgResp

func UpdateMemberStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusCodeUUIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member.NewUpdateMemberStatusLogic(r, svcCtx)
		resp, err := l.UpdateMemberStatus(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
