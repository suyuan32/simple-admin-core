package emailprovider

import (
	"context"

	"github.com/chimerakang/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/chimerakang/simple-admin-message-center/types/mcms"

	"github.com/chimerakang/simple-admin-core/api/internal/svc"
	"github.com/chimerakang/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmailProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmailProviderLogic {
	return &CreateEmailProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmailProviderLogic) CreateEmailProvider(req *types.EmailProviderInfo) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.CreateEmailProvider(l.ctx,
		&mcms.EmailProviderInfo{
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
