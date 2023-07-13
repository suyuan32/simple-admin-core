package emaillog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/emaillog"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /email_log/list emaillog GetEmailLogList
//
// Get email log list | 获取电子邮件日志列表
//
// Get email log list | 获取电子邮件日志列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: EmailLogListReq
//
// Responses:
//  200: EmailLogListResp

func GetEmailLogListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailLogListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emaillog.NewGetEmailLogListLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailLogList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
