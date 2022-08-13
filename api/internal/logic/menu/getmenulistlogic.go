package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList() (resp *types.MenuListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuByPage(context.Background(), &core.PageInfoReq{
		Page:     0,
		PageSize: 1000,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuListResp{}
	resp.Total = data.Total
	resp.Data = convertMenuList(data.Data)
	return resp, nil
}

func convertMenuList(data []*core.MenuInfo) []*types.Menu {
	if data == nil {
		return nil
	}
	var result []*types.Menu
	for _, v := range data {
		tmp := &types.Menu{
			BaseInfo: types.BaseInfo{
				ID:        uint(v.Id),
				CreatedAt: v.CreateAt,
				UpdatedAt: v.UpdateAt,
			},
			MenuType:  v.MenuType,
			ParentId:  uint(v.ParentId),
			MenuLevel: v.Level,
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			OrderNo:   v.OrderNo,
			Meta: types.Meta{
				KeepAlive:         v.Meta.KeepAlive,
				HideMenu:          v.Meta.HideMenu,
				HideBreadcrumb:    v.Meta.HideBreadcrumb,
				CurrentActiveMenu: v.Meta.CurrentActiveMenu,
				Title:             v.Meta.Title,
				Icon:              v.Meta.Icon,
				CloseTab:          v.Meta.CloseTab,
			},
			Children: convertMenuList(v.Children),
		}
		result = append(result, tmp)
	}
	return result
}
