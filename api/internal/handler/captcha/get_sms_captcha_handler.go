package captcha

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/captcha"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /captcha/sms captcha GetSmsCaptcha
//
// Get SMS Captcha | 获取短信验证码
//
// Get SMS Captcha | 获取短信验证码
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SmsCaptchaReq
//
// Responses:
//  200: BaseMsgResp

func GetSmsCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsCaptchaReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := captcha.NewGetSmsCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetSmsCaptcha(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
