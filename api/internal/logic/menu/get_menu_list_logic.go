package menu

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMenuListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMenuListLogic) GetMenuList() (resp *types.MenuPlainInfoListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuList(l.ctx, &core.PageInfoReq{
		Page:     1,
		PageSize: 100,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuPlainInfoListResp{}
	resp.Data.Total = data.Total
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.MenuPlainInfo{
			Id:                 v.Id,
			CreatedAt:          v.CreatedAt,
			UpdatedAt:          v.UpdatedAt,
			Trans:              l.svcCtx.Trans.Trans(l.lang, v.Meta.Title),
			MenuType:           v.MenuType,
			Level:              v.Level,
			Path:               v.Path,
			Name:               v.Name,
			Redirect:           v.Redirect,
			Component:          v.Component,
			Sort:               v.Sort,
			ParentId:           v.ParentId,
			Title:              v.Meta.Title,
			Icon:               v.Meta.Icon,
			HideMenu:           v.Meta.HideMenu,
			HideBreadcrumb:     v.Meta.HideBreadcrumb,
			IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
			HideTab:            v.Meta.HideTab,
			FrameSrc:           v.Meta.FrameSrc,
			CarryParam:         v.Meta.CarryParam,
			HideChildrenInMenu: v.Meta.HideChildrenInMenu,
			Affix:              v.Meta.Affix,
			DynamicLevel:       v.Meta.DynamicLevel,
			RealPath:           v.Meta.RealPath,
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	return resp, nil
}
