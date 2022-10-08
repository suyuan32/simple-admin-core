package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/oauth"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route POST /oauth/login oauth oauthLogin
// Oauth log in | Oauth 登录
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OauthLoginReq
// Responses:
//   200: RedirectResp
//   401: SimpleMsg
//   500: SimpleMsg

func OauthLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OauthLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := oauth.NewOauthLoginLogic(r.Context(), svcCtx)
		resp, err := l.OauthLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			//http.Redirect(w, r, resp.URL, http.StatusTemporaryRedirect)
			httpx.OkJson(w, &resp)
		}
	}
}
