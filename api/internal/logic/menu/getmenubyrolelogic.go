package menu

import (
	"context"
	"encoding/json"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuByRoleLogic {
	return &GetMenuByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuByRoleLogic) GetMenuByRole() (resp []*types.GetMenuListBase, err error) {
	roleId, _ := l.ctx.Value("roleId").(json.Number).Int64()
	data, err := l.svcCtx.CoreRpc.GetMenuListByRole(context.Background(), &core.IDReq{ID: uint64(roleId)})
	if err != nil {
		return nil, err
	}
	var result []*types.GetMenuListBase
	result = convertRoleMenuList(data.Data)
	return result, nil
}

func convertRoleMenuList(data []*core.MenuInfo) []*types.GetMenuListBase {
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
				KeepAlive:         v.Meta.KeepAlive,
				HideMenu:          v.Meta.HideMenu,
				HideBreadcrumb:    v.Meta.HideBreadcrumb,
				CurrentActiveMenu: v.Meta.CurrentActiveMenu,
				Title:             v.Meta.Title,
				Icon:              v.Meta.Icon,
				CloseTab:          v.Meta.CloseTab,
			},
			Children: convertRoleMenuList(v.Children),
		}
		result = append(result, tmp)
	}
	return result
}
