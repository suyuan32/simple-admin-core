package publicapi

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicSystemConfigurationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicSystemConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicSystemConfigurationListLogic {
	return &GetPublicSystemConfigurationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPublicSystemConfigurationListLogic) GetPublicSystemConfigurationList() (resp *types.ConfigurationListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetConfigurationList(l.ctx,
		&core.ConfigurationListReq{
			Page:     1,
			PageSize: 100,
			Category: pointy.GetPointer("system"),
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
