package logic

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTokenLogic {
	return &CreateOrUpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Token management
func (l *CreateOrUpdateTokenLogic) CreateOrUpdateToken(in *core.TokenInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.Token.Create().
			SetUUID(in.Uuid).
			SetToken(in.Token).
			SetStatus(uint8(in.Status)).
			SetSource(in.Source).
			SetExpiredAt(time.Unix(in.ExpiredAt, 0)).
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
		err := l.svcCtx.DB.Token.UpdateOneID(in.Id).
			SetUUID(in.Uuid).
			SetToken(in.Token).
			SetStatus(uint8(in.Status)).
			SetSource(in.Source).
			SetExpiredAt(time.Unix(in.ExpiredAt, 0)).
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
