package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.MenuPlainInfo) (resp *types.BaseMsgResp, err error) {
	if req.MenuType == 0 {
		req.Component = "LAYOUT"
		req.Path = ""
		req.Redirect = ""
		req.FrameSrc = ""
	}

	result, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        req.Id,
		MenuType:  req.MenuType,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Redirect:  req.Redirect,
		Component: req.Component,
		Sort:      req.Sort,
		Disabled:  req.Disabled,
		Meta: &core.Meta{
			Title:              req.Title,
			Icon:               req.Icon,
			HideMenu:           req.HideMenu,
			HideBreadcrumb:     req.HideBreadcrumb,
			IgnoreKeepAlive:    req.IgnoreKeepAlive,
			HideTab:            req.HideTab,
			FrameSrc:           req.FrameSrc,
			CarryParam:         req.CarryParam,
			HideChildrenInMenu: req.HideChildrenInMenu,
			Affix:              req.Affix,
			DynamicLevel:       req.DynamicLevel,
			RealPath:           req.RealPath,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
