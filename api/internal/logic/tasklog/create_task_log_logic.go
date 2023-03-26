package tasklog

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogLogic {
	return &CreateTaskLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogLogic) CreateTaskLog(req *types.TaskLogInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.JobRpc.CreateTaskLog(l.ctx,
		&job.TaskLogInfo{
			Id:         req.Id,
			StartedAt:  req.StartedAt,
			FinishedAt: req.FinishedAt,
			Result:     req.Result,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
