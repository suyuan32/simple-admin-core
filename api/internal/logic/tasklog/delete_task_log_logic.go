package tasklog

import (
	"context"

	"github.com/chimerakang/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/chimerakang/simple-admin-job/types/job"

	"github.com/chimerakang/simple-admin-core/api/internal/svc"
	"github.com/chimerakang/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogLogic {
	return &DeleteTaskLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTaskLogLogic) DeleteTaskLog(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	result, err := l.svcCtx.JobRpc.DeleteTaskLog(l.ctx, &job.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
