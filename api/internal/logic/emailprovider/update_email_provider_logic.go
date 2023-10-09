package emailprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailProviderLogic {
	return &UpdateEmailProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmailProviderLogic) UpdateEmailProvider(req *types.EmailProviderInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.McmsRpc.UpdateEmailProvider(l.ctx,
		&mcms.EmailProviderInfo{
			Id:        req.Id,
			Name:      req.Name,
			AuthType:  req.AuthType,
			EmailAddr: req.EmailAddr,
			Password:  req.Password,
			HostName:  req.HostName,
			Identify:  req.Identify,
			Secret:    req.Secret,
			Port:      req.Port,
			Tls:       req.Tls,
			IsDefault: req.IsDefault,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
