package menu

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMenuByRoleLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMenuByRoleLogic {
	return &GetMenuByRoleLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMenuByRoleLogic) GetMenuByRole() (resp *types.GetMenuListBaseResp, err error) {
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	data, err := l.svcCtx.CoreRpc.GetMenuListByRole(l.ctx, &core.IDReq{Id: uint64(roleId)})
	if err != nil {
		return nil, err
	}
	resp = &types.GetMenuListBaseResp{}
	resp.Data.Data = l.convertRoleMenuList(data.Data, l.lang)
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	return resp, nil
}

func (l *GetMenuByRoleLogic) convertRoleMenuList(data []*core.MenuInfo, lang string) []*types.GetMenuListBase {
	if data == nil {
		return nil
	}
	var result []*types.GetMenuListBase
	for _, v := range data {
		tmp := &types.GetMenuListBase{
			MenuType:  v.MenuType,
			MenuLevel: v.Level,
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			OrderNo:   v.OrderNo,
			Meta: types.Meta{
				Title:              l.svcCtx.Trans.Trans(lang, v.Meta.Title),
				Icon:               v.Meta.Icon,
				HideMenu:           v.Meta.HideMenu,
				HideBreadcrumb:     v.Meta.HideBreadcrumb,
				CurrentActiveMenu:  v.Meta.CurrentActiveMenu,
				IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
				HideTab:            v.Meta.HideTab,
				FrameSrc:           v.Meta.FrameSrc,
				CarryParam:         v.Meta.CarryParam,
				HideChildrenInMenu: v.Meta.HideChildrenInMenu,
				Affix:              v.Meta.Affix,
				DynamicLevel:       v.Meta.DynamicLevel,
				RealPath:           v.Meta.RealPath,
			},
			Children: l.convertRoleMenuList(v.Children, lang),
		}
		result = append(result, tmp)
	}
	return result
}
