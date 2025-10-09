package tasklog

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/chimerakang/simple-admin-job/types/job"

	"github.com/chimerakang/simple-admin-core/api/internal/svc"
	"github.com/chimerakang/simple-admin-core/api/internal/types"

	"github.com/chimerakang/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogListLogic {
	return &GetTaskLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogListLogic) GetTaskLogList(req *types.TaskLogListReq) (resp *types.TaskLogListResp, err error) {
	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.JobRpc.GetTaskLogList(l.ctx,
		&job.TaskLogListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			TaskId:   req.TaskId,
			Result:   req.Result,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.TaskLogListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.TaskLogInfo{
				Id:         v.Id,
				StartedAt:  v.StartedAt,
				FinishedAt: v.FinishedAt,
				Result:     v.Result,
			})
	}
	return resp, nil
}
