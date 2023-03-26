package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.TaskInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.JobRpc.CreateTask(l.ctx,
		&job.TaskInfo{
			Id:             req.Id,
			Status:         req.Status,
			Name:           req.Name,
			TaskGroup:      req.TaskGroup,
			CronExpression: req.CronExpression,
			Pattern:        req.Pattern,
			Payload:        req.Payload,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
