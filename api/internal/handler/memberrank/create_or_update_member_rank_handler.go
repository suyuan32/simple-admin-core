package memberrank

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/memberrank"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /member_rank/create_or_update memberrank CreateOrUpdateMemberRank
//
// Create or update member rank information | 创建或更新会员等级
//
// Create or update member rank information | 创建或更新会员等级
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateMemberRankReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateMemberRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateMemberRankReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := memberrank.NewCreateOrUpdateMemberRankLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateMemberRank(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
