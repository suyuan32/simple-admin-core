package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProviderListLogic {
	return &GetProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProviderListLogic) GetProviderList(in *core.PageInfoReq) (*core.ProviderListResp, error) {
	providers, err := l.svcCtx.DB.OauthProvider.Query().Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(errorx.DatabaseError)
	}

	//var providers []model.OauthProvider
	resp := &core.ProviderListResp{}
	resp.Total = providers.PageDetails.Total

	for _, v := range providers.List {
		resp.Data = append(resp.Data, &core.ProviderInfo{
			Id:           v.ID,
			Name:         v.Name,
			ClientId:     v.ClientID,
			ClientSecret: v.ClientSecret,
			RedirectUrl:  v.RedirectURL,
			Scopes:       v.Scopes,
			AuthUrl:      v.AuthURL,
			TokenUrl:     v.TokenURL,
			AuthStyle:    v.AuthStyle,
			InfoUrl:      v.InfoURL,
			CreatedAt:    v.CreatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
