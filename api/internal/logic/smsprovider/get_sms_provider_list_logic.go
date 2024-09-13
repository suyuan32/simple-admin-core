package smsprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderListLogic {
	return &GetSmsProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsProviderListLogic) GetSmsProviderList(req *types.SmsProviderListReq) (resp *types.SmsProviderListResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetSmsProviderList(l.ctx,
		&mcms.SmsProviderListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.SmsProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.SmsProviderInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:      v.Name,
				SecretId:  v.SecretId,
				SecretKey: v.SecretKey,
				Region:    v.Region,
				IsDefault: v.IsDefault,
			})
	}
	return resp, nil
}
