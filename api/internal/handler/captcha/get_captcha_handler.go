package captcha

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/captcha"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /captcha captcha GetCaptcha
//
// Get captcha | 获取验证码
//
// Get captcha | 获取验证码
//
// Responses:
//  200: CaptchaResp

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewGetCaptchaLogic(r, svcCtx)
		resp, err := l.GetCaptcha()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
