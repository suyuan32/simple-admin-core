package menu

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/menu"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /menu/role menu GetMenuByRole
//
// Get role's menu list API | 获取角色菜单列表
//
// Get role's menu list API | 获取角色菜单列表
//
// Responses:
//  200: GetMenuListBase
//  401: SimpleMsg
//  500: SimpleMsg

func GetMenuByRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetMenuByRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuByRole()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
