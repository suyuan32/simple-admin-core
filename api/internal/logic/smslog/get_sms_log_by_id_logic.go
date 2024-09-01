package smslog

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogByIdLogic {
	return &GetSmsLogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsLogByIdLogic) GetSmsLogById(req *types.UUIDReq) (resp *types.SmsLogInfoResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetSmsLogById(l.ctx, &mcms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.SmsLogInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.SmsLogInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			PhoneNumber: data.PhoneNumber,
			Content:     data.Content,
			SendStatus:  data.SendStatus,
			Provider:    data.Provider,
		},
	}, nil
}
