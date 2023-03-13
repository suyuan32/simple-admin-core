package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type CreateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOauthProviderLogic {
	return &CreateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOauthProviderLogic) CreateOauthProvider(in *core.OauthProviderInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.OauthProvider.Create().
		SetName(in.Name).
		SetClientID(in.ClientId).
		SetClientSecret(in.ClientSecret).
		SetRedirectURL(in.RedirectUrl).
		SetScopes(in.Scopes).
		SetAuthURL(in.AuthUrl).
		SetTokenURL(in.TokenUrl).
		SetAuthStyle(in.AuthStyle).
		SetInfoURL(in.InfoUrl).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
