package configuration

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateConfigurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateConfigurationLogic {
	return &CreateConfigurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateConfigurationLogic) CreateConfiguration(req *types.ConfigurationInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateConfiguration(l.ctx,
		&core.ConfigurationInfo{
			Sort:     req.Sort,
			State:    req.State,
			Name:     req.Name,
			Key:      req.Key,
			Value:    req.Value,
			Category: req.Category,
			Remark:   req.Remark,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
