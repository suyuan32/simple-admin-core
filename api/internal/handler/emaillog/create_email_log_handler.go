package emaillog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/emaillog"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /email_log/create emaillog CreateEmailLog
//
// Create email log information | 创建电子邮件日志
//
// Create email log information | 创建电子邮件日志
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmailLogInfo
//
// Responses:
//  200: BaseMsgResp

func CreateEmailLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailLogInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emaillog.NewCreateEmailLogLogic(r.Context(), svcCtx)
		resp, err := l.CreateEmailLog(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
