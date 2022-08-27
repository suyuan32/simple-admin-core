package user

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route POST /user/profile user updateUserProfile
// Update user's profile | 更新用户个人信息
// Responses:
//   200: SimpleMsg
//   401: SimpleMsg
//   500: SimpleMsg

func UpdateUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUpdateUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
