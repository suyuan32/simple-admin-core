package stock_movement

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockMovementListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockMovementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockMovementListLogic {
	return &GetStockMovementListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockMovementListLogic) GetStockMovementList(req *types.StockMovementListReq) (resp *types.StockMovementListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetStockMovementList(l.ctx, &core.StockMovementListReq{
		Page:         req.Page,
		PageSize:     req.PageSize,
		ProductId:    req.ProductId,
		MovementType: req.MovementType,
		Reference:    req.Reference,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.StockMovementListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.Total

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.StockMovementInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			ProductId:       v.ProductId,
			FromWarehouseId: v.FromWarehouseId,
			ToWarehouseId:   v.ToWarehouseId,
			Quantity:        v.Quantity,
			MovementType:    v.MovementType,
			Reference:       v.Reference,
			Details:         v.Details,
		})
	}
	return resp, nil
}
