package api

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetApiList(l.ctx,
		&core.ApiListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Path:        req.Path,
			Description: req.Description,
			Method:      req.Method,
			ApiGroup:    req.Group,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ApiListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.ApiInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Path:        v.Path,
				Trans:       l.svcCtx.Trans.Trans(l.ctx, v.Description),
				Description: v.Description,
				Group:       v.ApiGroup,
				Method:      v.Method,
			})
	}
	return resp, nil
}
