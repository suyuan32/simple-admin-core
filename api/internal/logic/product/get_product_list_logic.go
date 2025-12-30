package product

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req *types.ProductListReq) (resp *types.ProductListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetProductList(l.ctx, &core.ProductListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Sku:      req.Sku,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ProductListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.Total

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.ProductInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Status:      v.Status,
			Name:        v.Name,
			Sku:         v.Sku,
			Description: v.Description,
			Price:       v.Price,
			Unit:        v.Unit,
		})
	}
	return resp, nil
}
