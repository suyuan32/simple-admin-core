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

type GetEmailProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderListLogic {
	return &GetEmailProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailProviderListLogic) GetEmailProviderList(req *types.EmailProviderListReq) (resp *types.EmailProviderListResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetEmailProviderList(l.ctx,
		&mcms.EmailProviderListReq{
			Page:      req.Page,
			PageSize:  req.PageSize,
			Name:      req.Name,
			EmailAddr: req.EmailAddr,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.EmailProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.EmailProviderInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:      v.Name,
				AuthType:  v.AuthType,
				EmailAddr: v.EmailAddr,
				Password:  v.Password,
				HostName:  v.HostName,
				Identify:  v.Identify,
				Secret:    v.Secret,
				Port:      v.Port,
				Tls:       v.Tls,
				IsDefault: v.IsDefault,
			})
	}
	return resp, nil
}
