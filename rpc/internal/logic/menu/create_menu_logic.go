package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMenuLogic) CreateMenu(in *core.MenuInfo) (*core.BaseIDResp, error) {
	// get parent level
	var menuLevel uint32
	if in.ParentId != common.DefaultParentId {
		m, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.ParentId)).First(l.ctx)
		if err != nil {
			return nil, errorhandler.DefaultEntError(l.Logger, err, in)
		}

		menuLevel = m.MenuLevel + 1
	} else {
		menuLevel = 1
	}

	result, err := l.svcCtx.DB.Menu.Create().
		SetMenuLevel(menuLevel).
		SetMenuType(in.MenuType).
		SetParentID(in.ParentId).
		SetPath(in.Path).
		SetName(in.Name).
		SetRedirect(in.Redirect).
		SetComponent(in.Component).
		SetSort(in.Sort).
		SetDisabled(in.Disabled).
		// meta
		SetTitle(in.Meta.Title).
		SetIcon(in.Meta.Icon).
		SetHideMenu(in.Meta.HideMenu).
		SetHideBreadcrumb(in.Meta.HideBreadcrumb).
		SetIgnoreKeepAlive(in.Meta.IgnoreKeepAlive).
		SetHideTab(in.Meta.HideTab).
		SetFrameSrc(in.Meta.FrameSrc).
		SetCarryParam(in.Meta.CarryParam).
		SetHideChildrenInMenu(in.Meta.HideChildrenInMenu).
		SetAffix(in.Meta.Affix).
		SetDynamicLevel(in.Meta.DynamicLevel).
		SetRealPath(in.Meta.RealPath).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
