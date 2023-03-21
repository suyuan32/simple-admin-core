package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type UpdateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOauthProviderLogic {
	return &UpdateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOauthProviderLogic) UpdateOauthProvider(in *core.OauthProviderInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.OauthProvider.UpdateOneID(in.Id).
		SetNotEmptyName(in.Name).
		SetNotEmptyClientID(in.ClientId).
		SetNotEmptyClientSecret(in.ClientSecret).
		SetNotEmptyRedirectURL(in.RedirectUrl).
		SetNotEmptyScopes(in.Scopes).
		SetNotEmptyAuthURL(in.AuthUrl).
		SetNotEmptyTokenURL(in.TokenUrl).
		SetNotEmptyAuthStyle(in.AuthStyle).
		SetNotEmptyInfoURL(in.InfoUrl).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	if _, ok := providerConfig[in.Name]; ok {
		delete(providerConfig, in.Name)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
