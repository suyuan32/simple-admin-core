package warehouse

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWarehouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWarehouseLogic {
	return &UpdateWarehouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWarehouseLogic) UpdateWarehouse(req *types.WarehouseInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateWarehouse(l.ctx, &core.WarehouseInfo{
		Id:          req.Id,
		Status:      req.Status,
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
