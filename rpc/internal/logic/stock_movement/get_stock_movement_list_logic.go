package stock_movement

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/ent/stockmovement"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetStockMovementListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockMovementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockMovementListLogic {
	return &GetStockMovementListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockMovementListLogic) GetStockMovementList(in *core.StockMovementListReq) (*core.StockMovementListResp, error) {
	var predicates []predicate.StockMovement
	if in.CreatedAt != nil {
		predicates = append(predicates, stockmovement.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, stockmovement.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.DeletedAt != nil {
		predicates = append(predicates, stockmovement.DeletedAtGTE(time.UnixMilli(*in.DeletedAt)))
	}
	if in.ProductId != nil {
		predicates = append(predicates, stockmovement.ProductIDEQ(uuidx.ParseUUIDString(*in.ProductId)))
	}
	if in.FromWarehouseId != nil {
		predicates = append(predicates, stockmovement.FromWarehouseIDEQ(uuidx.ParseUUIDString(*in.FromWarehouseId)))
	}
	if in.ToWarehouseId != nil {
		predicates = append(predicates, stockmovement.ToWarehouseIDEQ(uuidx.ParseUUIDString(*in.ToWarehouseId)))
	}
	if in.Quantity != nil {
		predicates = append(predicates, stockmovement.QuantityEQ(*in.Quantity))
	}
	if in.MovementType != nil {
		predicates = append(predicates, stockmovement.MovementTypeContains(*in.MovementType))
	}
	if in.Reference != nil {
		predicates = append(predicates, stockmovement.ReferenceContains(*in.Reference))
	}
	if in.Details != nil {
		predicates = append(predicates, stockmovement.DetailsContains(*in.Details))
	}
	result, err := l.svcCtx.DB.StockMovement.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.StockMovementListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.StockMovementInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			ProductId:	pointy.GetPointer(v.ProductID.String()),
			FromWarehouseId:	pointy.GetPointer(v.FromWarehouseID.String()),
			ToWarehouseId:	pointy.GetPointer(v.ToWarehouseID.String()),
			Quantity:	&v.Quantity,
			MovementType:	&v.MovementType,
			Reference:	&v.Reference,
			Details:	&v.Details,
		})
	}

	return resp, nil
}
