package stock_movement

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockMovementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockMovementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockMovementLogic {
	return &UpdateStockMovementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockMovementLogic) UpdateStockMovement(in *core.StockMovementInfo) (*core.BaseResp, error) {
	err:= l.svcCtx.DB.StockMovement.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilProductID(uuidx.ParseUUIDStringToPointer(in.ProductId)).
			SetNotNilFromWarehouseID(uuidx.ParseUUIDStringToPointer(in.FromWarehouseId)).
			SetNotNilToWarehouseID(uuidx.ParseUUIDStringToPointer(in.ToWarehouseId)).
			SetNotNilQuantity(in.Quantity).
			SetNotNilMovementType(in.MovementType).
			SetNotNilReference(in.Reference).
			SetNotNilDetails(in.Details).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
