package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetMenuByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuByPageLogic {
	return &GetMenuByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuByPageLogic) GetMenuByPage(in *core.PageInfoReq) (*core.MenuInfoList, error) {
	var data []model.Menu
	result := l.svcCtx.DB.Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_no ASC")
	}).Limit(int(in.PageSize)).Offset(int(in.Page * in.PageSize)).
		Order("order_no ASC").Find(&data)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, "database error")
	}
	var res *core.MenuInfoList
	res = &core.MenuInfoList{}
	res.Data = findMenuChildren(data)
	// delete menus whose menu levels are not 1
	var tmp []*core.MenuInfo
	for _, v := range res.Data {
		if v.Level == 1 {
			tmp = append(tmp, v)
		}
	}
	res.Data = tmp
	res.Total = uint64(len(tmp))
	return res, nil
}

func findMenuChildren(data []model.Menu) []*core.MenuInfo {
	if data == nil {
		return nil
	}
	var result []*core.MenuInfo
	for _, v := range data {
		tmp := &core.MenuInfo{
			Id:        uint64(v.ID),
			CreateAt:  v.CreatedAt.UnixMilli(),
			UpdateAt:  v.UpdatedAt.UnixMilli(),
			MenuType:  v.MenuType,
			Level:     v.MenuLevel,
			ParentId:  uint32(v.ParentId),
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			OrderNo:   v.OrderNo,
			Meta: &core.Meta{
				KeepAlive:         v.Meta.KeepAlive,
				HideMenu:          v.Meta.HideMenu,
				HideBreadcrumb:    v.Meta.HideBreadcrumb,
				CurrentActiveMenu: v.Meta.CurrentActiveMenu,
				Title:             v.Meta.Title,
				Icon:              v.Meta.Icon,
				CloseTab:          v.Meta.CloseTab,
			},
			Children: findMenuChildren(v.Children),
		}
		result = append(result, tmp)
	}
	return result
}
