package tasklog

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-common/i18n"
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
