package stock_movement

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockMovementByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockMovementByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockMovementByIdLogic {
	return &GetStockMovementByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockMovementByIdLogic) GetStockMovementById(in *core.UUIDReq) (*core.StockMovementInfo, error) {
	result, err := l.svcCtx.DB.StockMovement.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.StockMovementInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		ProductId:	pointy.GetPointer(result.ProductID.String()),
		FromWarehouseId:	pointy.GetPointer(result.FromWarehouseID.String()),
		ToWarehouseId:	pointy.GetPointer(result.ToWarehouseID.String()),
		Quantity:	&result.Quantity,
		MovementType:	&result.MovementType,
		Reference:	&result.Reference,
		Details:	&result.Details,
	}, nil
}

