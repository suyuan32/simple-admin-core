package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskListLogic {
	return &GetTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskListLogic) GetTaskList(req *types.TaskListReq) (resp *types.TaskListResp, err error) {
	data, err := l.svcCtx.JobRpc.GetTaskList(l.ctx,
		&job.TaskListReq{
			Page:      req.Page,
			PageSize:  req.PageSize,
			Name:      req.Name,
			TaskGroup: req.TaskGroup,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.TaskListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.TaskInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:         v.Status,
				Name:           v.Name,
				TaskGroup:      v.TaskGroup,
				CronExpression: v.CronExpression,
				Pattern:        v.Pattern,
				Payload:        v.Payload,
			})
	}
	return resp, nil
}
