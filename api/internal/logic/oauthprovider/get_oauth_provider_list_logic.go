package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetOauthProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOauthProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(req *types.OauthProviderListReq) (resp *types.OauthProviderListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetOauthProviderList(l.ctx,
		&core.OauthProviderListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.OauthProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.OauthProviderInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:         v.Name,
				ClientId:     v.ClientId,
				ClientSecret: v.ClientSecret,
				RedirectUrl:  v.RedirectUrl,
				Scopes:       v.Scopes,
				AuthUrl:      v.AuthUrl,
				TokenUrl:     v.TokenUrl,
				AuthStyle:    v.AuthStyle,
				InfoUrl:      v.InfoUrl,
			})
	}
	return resp, nil
}
