package inventory

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInventoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInventoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInventoryListLogic {
	return &GetInventoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInventoryListLogic) GetInventoryList(req *types.InventoryListReq) (resp *types.InventoryListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetInventoryList(l.ctx, &core.InventoryListReq{
		Page:        req.Page,
		PageSize:    req.PageSize,
		ProductId:   req.ProductId,
		WarehouseId: req.WarehouseId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.InventoryListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.Total

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.InventoryInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			ProductId:   v.ProductId,
			WarehouseId: v.WarehouseId,
			Quantity:    v.Quantity,
		})
	}
	return resp, nil
}
