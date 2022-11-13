package menu

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateMenuLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuLogic {
	return &CreateOrUpdateMenuLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateMenuLogic) CreateOrUpdateMenu(req *types.CreateOrUpdateMenuReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateMenu(l.ctx, &core.CreateOrUpdateMenuReq{
		Id:        req.Id,
		MenuType:  req.MenuType,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Redirect:  req.Redirect,
		Component: req.Component,
		OrderNo:   req.OrderNo,
		Disabled:  req.Disabled,
		Meta: &core.Meta{
			Title:              req.Meta.Title,
			Icon:               req.Meta.Icon,
			HideMenu:           req.Meta.HideMenu,
			HideBreadcrumb:     req.Meta.HideBreadcrumb,
			CurrentActiveMenu:  req.Meta.CurrentActiveMenu,
			IgnoreKeepAlive:    req.Meta.IgnoreKeepAlive,
			HideTab:            req.Meta.HideTab,
			FrameSrc:           req.Meta.FrameSrc,
			CarryParam:         req.Meta.CarryParam,
			HideChildrenInMenu: req.Meta.HideChildrenInMenu,
			Affix:              req.Meta.Affix,
			DynamicLevel:       req.Meta.DynamicLevel,
			RealPath:           req.Meta.RealPath,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
