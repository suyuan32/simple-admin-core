package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/ent/role"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole(in *core.IDReq) (*core.MenuInfoList, error) {
	menus, err := l.svcCtx.DB.Role.Query().Where(role.ID(in.Id)).
		QueryMenus().Order(ent.Asc(menu.FieldOrderNo)).All(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	resp := &core.MenuInfoList{}
	resp.Total = uint64(len(menus))

	resp.Data = findRoleMenuChildren(menus, 1)

	return resp, nil

}

func findRoleMenuChildren(data []*ent.Menu, parentId uint64) []*core.MenuInfo {
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
				Children: findRoleMenuChildren(data, v.ID),
			}
			result = append(result, tmp)
		}

	}
	return result
}
