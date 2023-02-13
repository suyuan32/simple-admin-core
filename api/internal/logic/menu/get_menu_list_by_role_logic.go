package menu

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMenuListByRoleLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole() (resp *types.MenuListResp, err error) {
	roleId, _ := l.ctx.Value("roleId").(string)
	data, err := l.svcCtx.CoreRpc.GetMenuListByRole(l.ctx, &core.UUIDReq{Id: roleId})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuListResp{}
	resp.Data.Total = data.Total
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.MenuInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: 0,
				UpdatedAt: 0,
			},
			MenuType:  v.MenuType,
			Level:     v.Level,
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			Sort:      v.Sort,
			ParentId:  v.ParentId,
			Meta: types.Meta{
				Title:              l.svcCtx.Trans.Trans(l.lang, v.Meta.Title),
				Icon:               v.Meta.Icon,
				HideMenu:           v.Meta.HideMenu,
				HideBreadcrumb:     v.Meta.HideBreadcrumb,
				IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
				HideTab:            v.Meta.HideTab,
				FrameSrc:           v.Meta.FrameSrc,
				CarryParam:         v.Meta.CarryParam,
				HideChildrenInMenu: v.Meta.HideChildrenInMenu,
				Affix:              v.Meta.Affix,
				DynamicLevel:       v.Meta.DynamicLevel,
				RealPath:           v.Meta.RealPath,
			},
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	return resp, nil
}
