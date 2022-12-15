package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiLogic {
	return &CreateOrUpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateApiLogic) CreateOrUpdateApi(in *core.ApiInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.API.Create().
			SetPath(in.Path).
			SetDescription(in.Description).
			SetAPIGroup(in.Group).
			SetMethod(in.Method).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: i18n.CreateSuccess}, nil
	} else {
		err := l.svcCtx.DB.API.UpdateOneID(in.Id).
			SetPath(in.Path).
			SetDescription(in.Description).
			SetAPIGroup(in.Group).
			SetMethod(in.Method).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
	}
}
