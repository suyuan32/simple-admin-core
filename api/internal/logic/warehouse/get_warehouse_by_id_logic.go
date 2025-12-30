package warehouse

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWarehouseByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWarehouseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWarehouseByIdLogic {
	return &GetWarehouseByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWarehouseByIdLogic) GetWarehouseById(req *types.UUIDReq) (resp *types.WarehouseInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetWarehouseById(l.ctx, &core.UUIDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.WarehouseInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.WarehouseInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:      data.Status,
			Name:        data.Name,
			Location:    data.Location,
			Description: data.Description,
		},
	}, nil
}
