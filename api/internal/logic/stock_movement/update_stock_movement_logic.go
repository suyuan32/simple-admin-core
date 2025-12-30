package stock_movement

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockMovementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStockMovementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockMovementLogic {
	return &UpdateStockMovementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStockMovementLogic) UpdateStockMovement(req *types.StockMovementInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateStockMovement(l.ctx, &core.StockMovementInfo{
		Id:              req.Id,
		ProductId:       req.ProductId,
		FromWarehouseId: req.FromWarehouseId,
		ToWarehouseId:   req.ToWarehouseId,
		Quantity:        req.Quantity,
		MovementType:    req.MovementType,
		Reference:       req.Reference,
		Details:         req.Details,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
