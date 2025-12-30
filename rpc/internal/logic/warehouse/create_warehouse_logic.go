package warehouse

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWarehouseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWarehouseLogic {
	return &CreateWarehouseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWarehouseLogic) CreateWarehouse(in *core.WarehouseInfo) (*core.BaseUUIDResp, error) {
    query := l.svcCtx.DB.Warehouse.Create().
			SetNotNilName(in.Name).
			SetNotNilLocation(in.Location).
			SetNotNilDescription(in.Description)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
