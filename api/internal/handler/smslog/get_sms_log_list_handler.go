package smslog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/smslog"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sms_log/list smslog GetSmsLogList
//
// Get sms log list | 获取短信日志列表
//
// Get sms log list | 获取短信日志列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SmsLogListReq
//
// Responses:
//  200: SmsLogListResp

func GetSmsLogListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsLogListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := smslog.NewGetSmsLogListLogic(r.Context(), svcCtx)
		resp, err := l.GetSmsLogList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
