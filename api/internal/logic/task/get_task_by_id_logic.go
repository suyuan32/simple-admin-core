package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskByIdLogic {
	return &GetTaskByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskByIdLogic) GetTaskById(req *types.IDReq) (resp *types.TaskInfoResp, err error) {
	data, err := l.svcCtx.JobRpc.GetTaskById(l.ctx, &job.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.TaskInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TaskInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:         data.Status,
			Name:           data.Name,
			TaskGroup:      data.TaskGroup,
			CronExpression: data.CronExpression,
			Pattern:        data.Pattern,
			Payload:        data.Payload,
		},
	}, nil
}
