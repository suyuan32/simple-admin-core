package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *core.MenuInfo) (*core.BaseResp, error) {
	// get parent level
	var menuLevel uint32
	if *in.ParentId != common.DefaultParentId {
		m, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(*in.ParentId)).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		menuLevel = m.MenuLevel + 1
	} else {
		menuLevel = 1
	}

	err := l.svcCtx.DB.Menu.UpdateOneID(*in.Id).
		SetNotNilMenuLevel(&menuLevel).
		SetNotNilMenuType(in.MenuType).
		SetNotNilParentID(in.ParentId).
		SetNotNilPath(in.Path).
		SetNotNilName(in.Name).
		SetNotNilRedirect(in.Redirect).
		SetNotNilComponent(in.Component).
		SetNotNilSort(in.Sort).
		SetNotNilDisabled(in.Disabled).
		SetNotNilServiceName(in.ServiceName).
		SetNotNilPermission(in.Permission).
		// meta
		SetNotNilTitle(in.Meta.Title).
		SetNotNilIcon(in.Meta.Icon).
		SetNotNilHideMenu(in.Meta.HideMenu).
		SetNotNilHideBreadcrumb(in.Meta.HideBreadcrumb).
		SetNotNilIgnoreKeepAlive(in.Meta.IgnoreKeepAlive).
		SetNotNilHideTab(in.Meta.HideTab).
		SetNotNilFrameSrc(in.Meta.FrameSrc).
		SetNotNilCarryParam(in.Meta.CarryParam).
		SetNotNilHideChildrenInMenu(in.Meta.HideChildrenInMenu).
		SetNotNilAffix(in.Meta.Affix).
		SetNotNilDynamicLevel(in.Meta.DynamicLevel).
		SetNotNilRealPath(in.Meta.RealPath).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
