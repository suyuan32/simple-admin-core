package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListLogic) GetMenuList(in *core.PageInfoReq) (*core.MenuInfoList, error) {
	menus, err := l.svcCtx.DB.Menu.Query().Order(ent.Asc(menu.FieldOrderNo)).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(errorx.DatabaseError)
	}

	var resp *core.MenuInfoList
	resp = &core.MenuInfoList{}
	resp.Data = findMenuChildren(menus.List, uint64(1))
	resp.Total = uint64(len(resp.Data))
	return resp, nil
}

func findMenuChildren(data []*ent.Menu, parentId uint64) []*core.MenuInfo {
	if data == nil {
		return nil
	}
	var result []*core.MenuInfo
	for _, v := range data {
		if v.ParentID == parentId && v.ID != v.ParentID {
			tmp := &core.MenuInfo{
				Id:        v.ID,
				CreatedAt: v.CreatedAt.UnixMilli(),
				UpdatedAt: v.UpdatedAt.UnixMilli(),
				MenuType:  v.MenuType,
				Level:     v.MenuLevel,
				ParentId:  v.ParentID,
				Path:      v.Path,
				Name:      v.Name,
				Redirect:  v.Redirect,
				Component: v.Component,
				OrderNo:   v.OrderNo,
				Meta: &core.Meta{
					Title:              v.Title,
					Icon:               v.Icon,
					HideMenu:           v.HideMenu,
					HideBreadcrumb:     v.HideBreadcrumb,
					CurrentActiveMenu:  v.CurrentActiveMenu,
					IgnoreKeepAlive:    v.IgnoreKeepAlive,
					HideTab:            v.HideTab,
					FrameSrc:           v.FrameSrc,
					CarryParam:         v.CarryParam,
					HideChildrenInMenu: v.HideChildrenInMenu,
					Affix:              v.Affix,
					DynamicLevel:       v.DynamicLevel,
					RealPath:           v.RealPath,
				},
				Children: findMenuChildren(data, v.ID),
			}
			result = append(result, tmp)
		}
	}
	return result
}
