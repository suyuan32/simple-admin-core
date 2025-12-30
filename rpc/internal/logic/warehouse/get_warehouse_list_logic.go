package warehouse

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/ent/warehouse"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetWarehouseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWarehouseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWarehouseListLogic {
	return &GetWarehouseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWarehouseListLogic) GetWarehouseList(in *core.WarehouseListReq) (*core.WarehouseListResp, error) {
	var predicates []predicate.Warehouse
	if in.CreatedAt != nil {
		predicates = append(predicates, warehouse.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, warehouse.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.Status != nil {
		predicates = append(predicates, warehouse.StatusEQ(uint8(*in.Status)))
	}
	if in.DeletedAt != nil {
		predicates = append(predicates, warehouse.DeletedAtGTE(time.UnixMilli(*in.DeletedAt)))
	}
	if in.Name != nil {
		predicates = append(predicates, warehouse.NameContains(*in.Name))
	}
	if in.Location != nil {
		predicates = append(predicates, warehouse.LocationContains(*in.Location))
	}
	if in.Description != nil {
		predicates = append(predicates, warehouse.DescriptionContains(*in.Description))
	}
	result, err := l.svcCtx.DB.Warehouse.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.WarehouseListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.WarehouseInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Name:	&v.Name,
			Location:	&v.Location,
			Description:	&v.Description,
		})
	}

	return resp, nil
}
