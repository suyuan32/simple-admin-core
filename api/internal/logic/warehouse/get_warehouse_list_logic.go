package warehouse

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWarehouseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWarehouseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWarehouseListLogic {
	return &GetWarehouseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWarehouseListLogic) GetWarehouseList(req *types.WarehouseListReq) (resp *types.WarehouseListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetWarehouseList(l.ctx, &core.WarehouseListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Location: req.Location,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.WarehouseListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.Total

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.WarehouseInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Status:      v.Status,
			Name:        v.Name,
			Location:    v.Location,
			Description: v.Description,
		})
	}
	return resp, nil
}
