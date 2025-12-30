package inventory

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInventoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInventoryLogic {
	return &CreateInventoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInventoryLogic) CreateInventory(req *types.InventoryInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateInventory(l.ctx, &core.InventoryInfo{
		ProductId:   req.ProductId,
		WarehouseId: req.WarehouseId,
		Quantity:    req.Quantity,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
