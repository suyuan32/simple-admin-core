package configuration

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationByIdLogic {
	return &GetConfigurationByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationByIdLogic) GetConfigurationById(req *types.IDReq) (resp *types.ConfigurationInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetConfigurationById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.ConfigurationInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ConfigurationInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Sort:     data.Sort,
			State:    data.State,
			Name:     data.Name,
			Key:      data.Key,
			Value:    data.Value,
			Category: data.Category,
			Remark:   data.Remark,
		},
	}, nil
}
