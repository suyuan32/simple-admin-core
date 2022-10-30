package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/common/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/model"
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
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, "database error")
	}
	var res *core.MenuInfoList
	res = &core.MenuInfoList{}
	res.Data = findMenuChildren(data, 1)
	res.Total = uint64(len(res.Data))
	return res, nil
}

func findMenuChildren(data []model.Menu, parentId uint) []*core.MenuInfo {
	if data == nil {
		return nil
	}
	var result []*core.MenuInfo
	for _, v := range data {
		if v.ParentId == parentId && v.ID != v.ParentId {
			tmp := &core.MenuInfo{
				Id:        uint64(v.ID),
				CreatedAt: v.CreatedAt.UnixMilli(),
				UpdatedAt: v.UpdatedAt.UnixMilli(),
				MenuType:  v.MenuType,
				Level:     v.MenuLevel,
				ParentId:  uint32(v.ParentId),
				Path:      v.Path,
				Name:      v.Name,
				Redirect:  v.Redirect,
				Component: v.Component,
				OrderNo:   v.OrderNo,
				Meta: &core.Meta{
					Title:              v.Meta.Title,
					Icon:               v.Meta.Icon,
					HideMenu:           v.Meta.HideMenu,
					HideBreadcrumb:     v.Meta.HideBreadcrumb,
					CurrentActiveMenu:  v.Meta.CurrentActiveMenu,
					IgnoreKeepAlive:    v.Meta.IgnoreKeepAlive,
					HideTab:            v.Meta.HideTab,
					FrameSrc:           v.Meta.FrameSrc,
					CarryParam:         v.Meta.CarryParam,
					HideChildrenInMenu: v.Meta.HideChildrenInMenu,
					Affix:              v.Meta.Affix,
					DynamicLevel:       v.Meta.DynamicLevel,
					RealPath:           v.Meta.RealPath,
				},
				Children: findMenuChildren(data, v.ID),
			}
			result = append(result, tmp)
		}
	}
	return result
}
