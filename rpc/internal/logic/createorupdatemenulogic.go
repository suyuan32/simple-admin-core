package logic

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

//  menu service
func (l *CreateOrUpdateMenuLogic) CreateOrUpdateMenu(in *core.CreateOrUpdateMenuReq) (*core.BaseResp, error) {
	// get parent level
	var menuLevel uint32
	if in.ParentId != 0 {
		var parent model.Menu
		result := l.svcCtx.DB.Where("id = ?", in.ParentId).First(&parent)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, "parent menu is not exist")
		}
		menuLevel = parent.MenuLevel + 1
	} else {
		menuLevel = 1
	}
	var data *model.Menu
	if in.Id == 0 {
		// create menu
		data = &model.Menu{
			Model:     gorm.Model{},
			MenuType:  in.MenuType,
			MenuLevel: menuLevel,
			ParentId:  uint(in.ParentId),
			Path:      in.Path,
			Name:      in.Name,
			Redirect:  in.Redirect,
			Component: in.Component,
			OrderNo:   in.OrderNo,
			Disabled:  in.Disabled,
			Meta: model.Meta{
				KeepAlive:         in.Meta.KeepAlive,
				HideMenu:          in.Meta.HideMenu,
				HideBreadcrumb:    in.Meta.HideBreadcrumb,
				CurrentActiveMenu: in.Meta.CurrentActiveMenu,
				Title:             in.Meta.Title,
				Icon:              in.Meta.Icon,
				CloseTab:          in.Meta.CloseTab,
			},
		}
		result := l.svcCtx.DB.Create(data)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, "menu exist")
		}
		return &core.BaseResp{Msg: "common.createSuccess"}, nil
	} else {
		var origin *model.Menu
		result := l.svcCtx.DB.Where("id = ?", in.Id).First(&origin)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, "sys.menu.menuNotExists")
		}
		data = &model.Menu{
			Model:     gorm.Model{ID: uint(in.Id), CreatedAt: origin.CreatedAt, UpdatedAt: time.Now()},
			MenuLevel: menuLevel,
			MenuType:  in.MenuType,
			ParentId:  uint(in.ParentId),
			Path:      in.Path,
			Name:      in.Name,
			Redirect:  in.Redirect,
			Component: in.Component,
			OrderNo:   in.OrderNo,
			Disabled:  in.Disabled,
			Meta: model.Meta{
				KeepAlive:         in.Meta.KeepAlive,
				HideMenu:          in.Meta.HideMenu,
				HideBreadcrumb:    in.Meta.HideBreadcrumb,
				CurrentActiveMenu: in.Meta.CurrentActiveMenu,
				Title:             in.Meta.Title,
				Icon:              in.Meta.Icon,
				CloseTab:          in.Meta.CloseTab,
			},
		}
		result = l.svcCtx.DB.Save(data)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			return nil, status.Error(codes.InvalidArgument, "sys.menu.menuNotExists")
		}
		return &core.BaseResp{Msg: "common.updateSuccess"}, nil
	}
}
