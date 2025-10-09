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

type GetTaskLogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogByIdLogic {
	return &GetTaskLogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogByIdLogic) GetTaskLogById(req *types.IDReq) (resp *types.TaskLogInfoResp, err error) {
	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.JobRpc.GetTaskLogById(l.ctx, &job.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.TaskLogInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TaskLogInfo{
			Id:         data.Id,
			StartedAt:  data.StartedAt,
			FinishedAt: data.FinishedAt,
			Result:     data.Result,
		},
	}, nil
}
