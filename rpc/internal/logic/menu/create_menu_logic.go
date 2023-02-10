package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/enum"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
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

func (l *CreateMenuLogic) CreateMenu(in *core.MenuInfo) (*core.BaseResp, error) {
	// get parent level
	var menuLevel uint32
	if in.ParentId != enum.DefaultParentId {
		m, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.ParentId)).First(l.ctx)
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

		menuLevel = m.MenuLevel + 1
	} else {
		menuLevel = 1
	}

	err := l.svcCtx.DB.Menu.Create().
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
		Exec(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &core.BaseResp{Msg: i18n.CreateSuccess}, nil
}
