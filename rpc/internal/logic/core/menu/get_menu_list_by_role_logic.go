package menu

import (
	"context"
	"strings"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
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

func (l *GetMenuListByRoleLogic) GetMenuListByRole(in *core.BaseMsg) (*core.MenuInfoList, error) {
	roles, err := l.svcCtx.DB.Role.Query().Where(role.CodeIn(strings.Split(in.Msg, ",")...)).WithMenus(func(query *ent.MenuQuery) {
		query.Order(ent.Asc(menu.FieldSort))
		query.Where(menu.Disabled(false))
	}).All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MenuInfoList{}

	existMap := map[uint64]struct{}{}
	for _, r := range roles {
		for _, m := range r.Edges.Menus {
			if _, ok := existMap[m.ID]; !ok {
				resp.Data = append(resp.Data, &core.MenuInfo{
					Id:        &m.ID,
					CreatedAt: pointy.GetPointer(m.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(m.UpdatedAt.UnixMilli()),
					MenuType:  &m.MenuType,
					Level:     &m.MenuLevel,
					ParentId:  &m.ParentID,
					Path:      &m.Path,
					Name:      &m.Name,
					Redirect:  &m.Redirect,
					Component: &m.Component,
					Sort:      &m.Sort,
					Meta: &core.Meta{
						Title:              &m.Title,
						Icon:               &m.Icon,
						HideMenu:           &m.HideMenu,
						HideBreadcrumb:     &m.HideBreadcrumb,
						IgnoreKeepAlive:    &m.IgnoreKeepAlive,
						HideTab:            &m.HideTab,
						FrameSrc:           &m.FrameSrc,
						CarryParam:         &m.CarryParam,
						HideChildrenInMenu: &m.HideChildrenInMenu,
						Affix:              &m.Affix,
						DynamicLevel:       &m.DynamicLevel,
						RealPath:           &m.RealPath,
					},
				})
				existMap[m.ID] = struct{}{}
			}
		}
	}

	resp.Total = uint64(len(resp.Data))

	return resp, nil
}
