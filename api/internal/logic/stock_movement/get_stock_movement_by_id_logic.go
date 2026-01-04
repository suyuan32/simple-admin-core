package stock_movement

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockMovementByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockMovementByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockMovementByIdLogic {
	return &GetStockMovementByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockMovementByIdLogic) GetStockMovementById(req *types.UUIDReq) (resp *types.StockMovementInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetStockMovementById(l.ctx, &core.UUIDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.StockMovementInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StockMovementInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			ProductId:       data.ProductId,
			FromWarehouseId: data.FromWarehouseId,
			ToWarehouseId:   data.ToWarehouseId,
			Quantity:        data.Quantity,
			MovementType:    data.MovementType,
			Reference:       data.Reference,
			Details:         data.Details,
		},
	}, nil
}
