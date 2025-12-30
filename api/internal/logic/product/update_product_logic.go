package product

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	core "github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.ProductInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateProduct(l.ctx, &core.ProductInfo{
		Id:          req.Id,
		Status:      req.Status,
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		Unit:        req.Unit,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
