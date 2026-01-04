package inventory

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/ent/inventory"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetInventoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInventoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInventoryListLogic {
	return &GetInventoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInventoryListLogic) GetInventoryList(in *core.InventoryListReq) (*core.InventoryListResp, error) {
	var predicates []predicate.Inventory
	if in.CreatedAt != nil {
		predicates = append(predicates, inventory.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, inventory.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.DeletedAt != nil {
		predicates = append(predicates, inventory.DeletedAtGTE(time.UnixMilli(*in.DeletedAt)))
	}
	if in.ProductId != nil {
		predicates = append(predicates, inventory.ProductIDEQ(uuidx.ParseUUIDString(*in.ProductId)))
	}
	if in.WarehouseId != nil {
		predicates = append(predicates, inventory.WarehouseIDEQ(uuidx.ParseUUIDString(*in.WarehouseId)))
	}
	if in.Quantity != nil {
		predicates = append(predicates, inventory.QuantityEQ(*in.Quantity))
	}
	result, err := l.svcCtx.DB.Inventory.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.InventoryListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.InventoryInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			ProductId:	pointy.GetPointer(v.ProductID.String()),
			WarehouseId:	pointy.GetPointer(v.WarehouseID.String()),
			Quantity:	&v.Quantity,
		})
	}

	return resp, nil
}
