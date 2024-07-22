package configuration

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationListLogic {
	return &GetConfigurationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationListLogic) GetConfigurationList(req *types.ConfigurationListReq) (resp *types.ConfigurationListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetConfigurationList(l.ctx,
		&core.ConfigurationListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
			Key:      req.Key,
			Category: req.Category,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ConfigurationListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.ConfigurationInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Sort:     v.Sort,
				State:    v.State,
				Name:     v.Name,
				Key:      v.Key,
				Value:    v.Value,
				Category: v.Category,
				Remark:   v.Remark,
			})
	}
	return resp, nil
}
