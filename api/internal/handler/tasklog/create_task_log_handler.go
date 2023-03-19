package tasklog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/tasklog"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /task_log/create tasklog CreateTaskLog
//
// Create task log information | 创建任务日志
//
// Create task log information | 创建任务日志
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TaskLogInfo
//
// Responses:
//  200: BaseMsgResp

func CreateTaskLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskLogInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tasklog.NewCreateTaskLogLogic(r.Context(), svcCtx)
		resp, err := l.CreateTaskLog(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
