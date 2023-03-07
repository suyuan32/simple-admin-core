package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
)

// swagger:route get /user/profile user GetUserProfile
//
// Get user's profile | 获取用户个人信息
//
// Get user's profile | 获取用户个人信息
//
// Responses:
//  200: ProfileResp

func GetUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetUserProfile()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
