package warehouse

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWarehouseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWarehouseLogic {
	return &UpdateWarehouseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWarehouseLogic) UpdateWarehouse(in *core.WarehouseInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.Warehouse.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilName(in.Name).
			SetNotNilLocation(in.Location).
			SetNotNilDescription(in.Description)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
