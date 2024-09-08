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

type UpdateEmailLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmailLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogLogic {
	return &UpdateEmailLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmailLogLogic) UpdateEmailLog(req *types.EmailLogInfo) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.UpdateEmailLog(l.ctx,
		&mcms.EmailLogInfo{
			Id:         req.Id,
			Target:     req.Target,
			Subject:    req.Subject,
			Content:    req.Content,
			SendStatus: req.SendStatus,
			Provider:   req.Provider,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
