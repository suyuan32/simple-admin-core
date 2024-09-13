package emailprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderByIdLogic {
	return &GetEmailProviderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailProviderByIdLogic) GetEmailProviderById(req *types.IDReq) (resp *types.EmailProviderInfoResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetEmailProviderById(l.ctx, &mcms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.EmailProviderInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.EmailProviderInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:      data.Name,
			AuthType:  data.AuthType,
			EmailAddr: data.EmailAddr,
			Password:  data.Password,
			HostName:  data.HostName,
			Identify:  data.Identify,
			Secret:    data.Secret,
			Port:      data.Port,
			Tls:       data.Tls,
			IsDefault: data.IsDefault,
		},
	}, nil
}
