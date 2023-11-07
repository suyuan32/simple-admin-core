package base

import (
	"context"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

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
	if !l.svcCtx.Config.ProjectConf.AllowInit {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.PermissionDeny)
	}

	if !l.svcCtx.Config.JobRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	result, err := l.svcCtx.JobRpc.InitDatabase(l.ctx, &job.Empty{})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin Policy", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.DatabaseError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
