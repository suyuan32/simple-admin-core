package base

import (
	"context"

	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitJobDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitJobDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitJobDatabaseLogic {
	return &InitJobDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitJobDatabaseLogic) InitJobDatabase() (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.JobRpc.InitDatabase(l.ctx, &job.Empty{})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
