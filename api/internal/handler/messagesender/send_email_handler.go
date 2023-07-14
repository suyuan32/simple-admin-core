package messagesender

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/messagesender"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /email/send messagesender SendEmail
//
// Send email message | 发送电子邮件
//
// Send email message | 发送电子邮件
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SendEmailReq
//
// Responses:
//  200: BaseMsgResp

func SendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendEmailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := messagesender.NewSendEmailLogic(r.Context(), svcCtx)
		resp, err := l.SendEmail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
