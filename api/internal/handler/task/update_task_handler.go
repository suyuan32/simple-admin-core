package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/task"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /task/update task UpdateTask
//
// Update task information | 更新Task
//
// Update task information | 更新Task
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TaskInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewUpdateTaskLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTask(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
