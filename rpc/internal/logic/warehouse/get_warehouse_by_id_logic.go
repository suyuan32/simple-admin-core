package warehouse

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetWarehouseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWarehouseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWarehouseByIdLogic {
	return &GetWarehouseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWarehouseByIdLogic) GetWarehouseById(in *core.UUIDReq) (*core.WarehouseInfo, error) {
	result, err := l.svcCtx.DB.Warehouse.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.WarehouseInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		Name:	&result.Name,
		Location:	&result.Location,
		Description:	&result.Description,
	}, nil
}

