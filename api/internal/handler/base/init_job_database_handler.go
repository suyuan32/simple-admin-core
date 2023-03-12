package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/base"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
)

// swagger:route get /core/init/job_database base InitJobDatabase
//
// Initialize job database | 初始化定时任务数据库
//
// Initialize job database | 初始化定时任务数据库
//
// Responses:
//  200: BaseMsgResp

func InitJobDatabaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewInitJobDatabaseLogic(r.Context(), svcCtx)
		resp, err := l.InitJobDatabase()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
