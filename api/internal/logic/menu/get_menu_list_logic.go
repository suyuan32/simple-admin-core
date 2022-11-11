package menu

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetMenuListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetMenuListLogic) GetMenuList() (resp *types.MenuListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuList(l.ctx, &core.PageInfoReq{
		Page:     1,
		PageSize: 100,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuListResp{}
	resp.Total = data.Total
	resp.Data = l.convertMenuList(data.Data, l.r.Header.Get("Accept-Language"))
	return resp, nil
}

func (l *GetMenuListLogic) convertMenuList(data []*core.MenuInfo, lang string) []*types.MenuInfo {
	if data == nil {
		return nil
	}
	var result []*types.MenuInfo
	for _, v := range data {
		tmp := &types.MenuInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Trans:     l.svcCtx.Trans.Trans(lang, v.Meta.Title),
			MenuType:  v.MenuType,
			ParentId:  v.ParentId,
			MenuLevel: v.Level,
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			OrderNo:   v.OrderNo,
			Meta: types.Meta{
				Title:              v.Meta.Title,
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
			Children: l.convertMenuList(v.Children, lang),
		}
		result = append(result, tmp)
	}
	return result
}
