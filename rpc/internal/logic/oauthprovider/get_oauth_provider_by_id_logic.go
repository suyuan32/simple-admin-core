package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(in *core.IDReq) (*core.OauthProviderInfo, error) {
	result, err := l.svcCtx.DB.OauthProvider.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.OauthProviderInfo{
		Id:           result.ID,
		CreatedAt:    result.CreatedAt.UnixMilli(),
		UpdatedAt:    result.UpdatedAt.UnixMilli(),
		Name:         result.Name,
		ClientId:     result.ClientID,
		ClientSecret: result.ClientSecret,
		RedirectUrl:  result.RedirectURL,
		Scopes:       result.Scopes,
		AuthUrl:      result.AuthURL,
		TokenUrl:     result.TokenURL,
		AuthStyle:    result.AuthStyle,
		InfoUrl:      result.InfoURL,
	}, nil
}
