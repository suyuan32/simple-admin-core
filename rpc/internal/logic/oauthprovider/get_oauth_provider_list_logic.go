package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(in *core.OauthProviderListReq) (*core.OauthProviderListResp, error) {
	var predicates []predicate.OauthProvider
	if in.Name != "" {
		predicates = append(predicates, oauthprovider.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.OauthProvider.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.OauthProviderListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.OauthProviderInfo{
			Id:           v.ID,
			CreatedAt:    v.CreatedAt.UnixMilli(),
			UpdatedAt:    v.UpdatedAt.UnixMilli(),
			Name:         v.Name,
			ClientId:     v.ClientID,
			ClientSecret: v.ClientSecret,
			RedirectUrl:  v.RedirectURL,
			Scopes:       v.Scopes,
			AuthUrl:      v.AuthURL,
			TokenUrl:     v.TokenURL,
			AuthStyle:    v.AuthStyle,
			InfoUrl:      v.InfoURL,
		})
	}

	return resp, nil
}
