package oauthprovider

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type GetOauthProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetOauthProviderListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(req *types.OauthProviderListReq) (resp *types.OauthProviderListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetOauthProviderList(l.ctx,
		&core.OauthProviderListReq{
			Page:         req.Page,
			PageSize:     req.PageSize,
			Name:         req.Name,
			ClientId:     req.ClientId,
			ClientSecret: req.ClientSecret,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.OauthProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.OauthProviderInfo{
				BaseInfo: types.BaseInfo{
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
