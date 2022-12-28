package tenant

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTenantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTenantLogic {
	return &CreateOrUpdateTenantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateTenantLogic) CreateOrUpdateTenant(in *core.CreateOrUpdateTenantReq) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line
	// var tenantLevel uint32
	// if in.Pid != "" {
	// 	t, err := l.svcCtx.DB.Tenant.Query().Where(tenant.UUIDEQ(in.Pid)).First(l.ctx)
	// 	if err != nil {
	// 		switch {
	// 		case ent.IsNotFound(err):
	// 			logx.Errorw(err.Error(), logx.Field("detail", in))
	// 			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
	// 		default:
	// 			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
	// 			return nil, statuserr.NewInternalError(i18n.DatabaseError)
	// 		}
	// 	}
	// }

	return &core.BaseResp{}, nil
}
