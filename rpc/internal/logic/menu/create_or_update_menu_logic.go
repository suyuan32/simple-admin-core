package menu

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuLogic {
	return &CreateOrUpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateMenuLogic) CreateOrUpdateMenu(in *core.CreateOrUpdateMenuReq) (*core.BaseResp, error) {
	// get parent level
	var menuLevel uint32
	if in.ParentId != 0 {
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

	if in.Id == 0 {
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
			SetCurrentActiveMenu(in.Meta.CurrentActiveMenu).
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
	} else {
		if in.ParentId != 0 {
			exist, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.ParentId)).Exist(l.ctx)
			if err != nil {
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, err
			}

			if !exist {
				logx.Errorw("menu not found", logx.Field("menuId", in.Id))
				return nil, status.Error(codes.InvalidArgument, "menu.menuNotExists")
			}
		}

		err := l.svcCtx.DB.Menu.UpdateOneID(in.Id).
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
			SetCurrentActiveMenu(in.Meta.CurrentActiveMenu).
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
				return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
	}
}
