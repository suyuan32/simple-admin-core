package product

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductByIdLogic) GetProductById(req *types.UUIDReq) (resp *types.ProductInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetProductById(l.ctx, &core.UUIDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.ProductInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ProductInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:      data.Status,
			Name:        data.Name,
			Sku:         data.Sku,
			Description: data.Description,
			Price:       data.Price,
			Unit:        data.Unit,
		},
	}, nil
}
