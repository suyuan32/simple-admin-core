package inventory

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInventoryLogic {
	return &CreateInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateInventoryLogic) CreateInventory(in *core.InventoryInfo) (*core.BaseUUIDResp, error) {
    result, err := l.svcCtx.DB.Inventory.Create().
			SetNotNilProductID(uuidx.ParseUUIDStringToPointer(in.ProductId)).
			SetNotNilWarehouseID(uuidx.ParseUUIDStringToPointer(in.WarehouseId)).
			SetNotNilQuantity(in.Quantity).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
