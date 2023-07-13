package smsprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSmsProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSmsProviderLogic {
	return &CreateSmsProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSmsProviderLogic) CreateSmsProvider(req *types.SmsProviderInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.CreateSmsProvider(l.ctx,
		&mcms.SmsProviderInfo{
			Name:      req.Name,
			SecretId:  req.SecretId,
			SecretKey: req.SecretKey,
			Region:    req.Region,
			IsDefault: req.IsDefault,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
