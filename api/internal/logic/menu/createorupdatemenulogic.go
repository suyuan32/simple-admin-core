package menu

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuLogic {
	return &CreateOrUpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateMenuLogic) CreateOrUpdateMenu(req *types.CreateOrUpdateMenuReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateMenu(context.Background(), &core.CreateOrUpdateMenuReq{
		Id:        uint64(req.ID),
		MenuType:  req.MenuType,
		ParentId:  uint32(req.ParentId),
		Path:      req.Path,
		Name:      req.Name,
		Redirect:  req.Redirect,
		Component: req.Component,
		OrderNo:   req.OrderNo,
		Disabled:  req.Disabled,
		Meta: &core.Meta{
			KeepAlive:         req.KeepAlive,
			HideMenu:          req.HideMenu,
			HideBreadcrumb:    req.HideBreadcrumb,
			CurrentActiveMenu: req.CurrentActiveMenu,
			Title:             req.Title,
			Icon:              req.Icon,
			CloseTab:          req.CloseTab,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
