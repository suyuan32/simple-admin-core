package memberrank

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/memberrank"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /member_rank/update memberrank UpdateMemberRank
//
// Update member rank information | 更新MemberRank
//
// Update member rank information | 更新MemberRank
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MemberRankInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateMemberRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberRankInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := memberrank.NewUpdateMemberRankLogic(r, svcCtx)
		resp, err := l.UpdateMemberRank(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
