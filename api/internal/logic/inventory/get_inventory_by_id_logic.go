package inventory

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInventoryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInventoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInventoryByIdLogic {
	return &GetInventoryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInventoryByIdLogic) GetInventoryById(req *types.UUIDReq) (resp *types.InventoryInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetInventoryById(l.ctx, &core.UUIDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.InventoryInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.InventoryInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			ProductId:   data.ProductId,
			WarehouseId: data.WarehouseId,
			Quantity:    data.Quantity,
		},
	}, nil
}
