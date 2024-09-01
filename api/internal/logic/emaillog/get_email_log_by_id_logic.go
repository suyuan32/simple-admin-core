package emaillog

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailLogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailLogByIdLogic {
	return &GetEmailLogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailLogByIdLogic) GetEmailLogById(req *types.UUIDReq) (resp *types.EmailLogInfoResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetEmailLogById(l.ctx, &mcms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.EmailLogInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.EmailLogInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Target:     data.Target,
			Subject:    data.Subject,
			Content:    data.Content,
			SendStatus: data.SendStatus,
			Provider:   data.Provider,
		},
	}, nil
}
