package smslog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/smslog"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sms_log/create smslog CreateSmsLog
//
// Create sms log information | 创建短信日志
//
// Create sms log information | 创建短信日志
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SmsLogInfo
//
// Responses:
//  200: BaseMsgResp

func CreateSmsLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsLogInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := smslog.NewCreateSmsLogLogic(r.Context(), svcCtx)
		resp, err := l.CreateSmsLog(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
