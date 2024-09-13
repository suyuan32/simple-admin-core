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

type GetSmsProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderByIdLogic {
	return &GetSmsProviderByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsProviderByIdLogic) GetSmsProviderById(req *types.IDReq) (resp *types.SmsProviderInfoResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetSmsProviderById(l.ctx, &mcms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.SmsProviderInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.SmsProviderInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:      data.Name,
			SecretId:  data.SecretId,
			SecretKey: data.SecretKey,
			Region:    data.Region,
			IsDefault: data.IsDefault,
		},
	}, nil
}
