package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateProviderLogic {
	return &CreateOrUpdateProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// oauth management
func (l *CreateOrUpdateProviderLogic) CreateOrUpdateProvider(in *core.ProviderInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.OauthProvider.Create().
			SetName(in.Name).
			SetClientID(in.ClientId).
			SetClientSecret(in.ClientSecret).
			SetRedirectURL(in.RedirectUrl).
			SetScopes(in.Scopes).
			SetAuthURL(in.AuthUrl).
			SetTokenURL(in.TokenUrl).
			SetAuthStyle(in.AuthStyle).
			SetInfoURL(in.InfoUrl).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.CreateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		err := l.svcCtx.DB.OauthProvider.UpdateOneID(in.Id).
			SetName(in.Name).
			SetClientID(in.ClientId).
			SetClientSecret(in.ClientSecret).
			SetRedirectURL(in.RedirectUrl).
			SetScopes(in.Scopes).
			SetAuthURL(in.AuthUrl).
			SetTokenURL(in.TokenUrl).
			SetAuthStyle(in.AuthStyle).
			SetInfoURL(in.InfoUrl).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.UpdateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		delete(providerConfig, in.Name)

		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}
