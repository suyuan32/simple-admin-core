package product

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/rpc/ent/product"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListLogic) GetProductList(in *core.ProductListReq) (*core.ProductListResp, error) {
	var predicates []predicate.Product
	if in.CreatedAt != nil {
		predicates = append(predicates, product.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, product.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.Status != nil {
		predicates = append(predicates, product.StatusEQ(uint8(*in.Status)))
	}
	if in.DeletedAt != nil {
		predicates = append(predicates, product.DeletedAtGTE(time.UnixMilli(*in.DeletedAt)))
	}
	if in.Name != nil {
		predicates = append(predicates, product.NameContains(*in.Name))
	}
	if in.Sku != nil {
		predicates = append(predicates, product.SkuContains(*in.Sku))
	}
	if in.Description != nil {
		predicates = append(predicates, product.DescriptionContains(*in.Description))
	}
	if in.Price != nil {
		predicates = append(predicates, product.PriceEQ(*in.Price))
	}
	if in.Unit != nil {
		predicates = append(predicates, product.UnitContains(*in.Unit))
	}
	result, err := l.svcCtx.DB.Product.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.ProductListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.ProductInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Name:	&v.Name,
			Sku:	&v.Sku,
			Description:	&v.Description,
			Price:	&v.Price,
			Unit:	&v.Unit,
		})
	}

	return resp, nil
}
