package menu

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
		QueryMenus().Where().Order(ent.Asc(menu.FieldSort)).All(l.ctx)
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

	for _, v := range menus {
		resp.Data = append(resp.Data, &core.MenuInfo{
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
			Sort:      v.Sort,
			Meta: &core.Meta{
				Title:              v.Title,
				Icon:               v.Icon,
				HideMenu:           v.HideMenu,
				HideBreadcrumb:     v.HideBreadcrumb,
				IgnoreKeepAlive:    v.IgnoreKeepAlive,
				HideTab:            v.HideTab,
				FrameSrc:           v.FrameSrc,
				CarryParam:         v.CarryParam,
				HideChildrenInMenu: v.HideChildrenInMenu,
				Affix:              v.Affix,
				DynamicLevel:       v.DynamicLevel,
				RealPath:           v.RealPath,
			},
		})
	}

	return resp, nil
}
