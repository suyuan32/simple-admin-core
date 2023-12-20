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
			ServiceName: req.ServiceName,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ApiListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	// Translate api group
	var group = make(map[string]string)
	for _, v := range data.Data {
		// If page size is over 1000, use api group i18n translation. Mainly use in authority list.
		if req.PageSize > 1000 {
			if _, exist := group[*v.ApiGroup]; !exist {
				group[*v.ApiGroup] = l.svcCtx.Trans.Trans(l.ctx, "apiGroup."+*v.ApiGroup)
			}
			*v.ApiGroup = group[*v.ApiGroup]
		}
		resp.Data.Data = append(resp.Data.Data,
			types.ApiInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Path:        v.Path,
				Trans:       l.svcCtx.Trans.Trans(l.ctx, *v.Description),
				Description: v.Description,
				Group:       v.ApiGroup,
				Method:      v.Method,
				IsRequired:  v.IsRequired,
				ServiceName: v.ServiceName,
			})
	}
	return resp, nil
}
