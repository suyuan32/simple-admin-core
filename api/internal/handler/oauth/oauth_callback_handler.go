package oauth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/oauth"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
)

// swagger:route get /oauth/login/callback oauth OauthCallback
//
// Oauth log in callback route | Oauth 登录返回调接口
//
// Oauth log in callback route | Oauth 登录返回调接口
//
// Responses:
//  200: CallbackResp
//  401: SimpleMsg
//  500: SimpleMsg

func OauthCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewOauthCallbackLogic(r, svcCtx)
		resp, err := l.OauthCallback()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
