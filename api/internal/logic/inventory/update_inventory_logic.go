package inventory

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInventoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInventoryLogic {
	return &UpdateInventoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInventoryLogic) UpdateInventory(req *types.InventoryInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateInventory(l.ctx, &core.InventoryInfo{
		Id:          req.Id,
		ProductId:   req.ProductId,
		WarehouseId: req.WarehouseId,
		Quantity:    req.Quantity,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
