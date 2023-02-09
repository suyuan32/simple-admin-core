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

type GetOauthProviderByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetOauthProviderByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(req *types.IDReq) (resp *types.OauthProviderInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetOauthProviderById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.OauthProviderInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.OauthProviderInfo{
			BaseInfo: types.BaseInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:         data.Name,
			ClientId:     data.ClientId,
			ClientSecret: data.ClientSecret,
			RedirectUrl:  data.RedirectUrl,
			Scopes:       data.Scopes,
			AuthUrl:      data.AuthUrl,
			TokenUrl:     data.TokenUrl,
			AuthStyle:    data.AuthStyle,
			InfoUrl:      data.InfoUrl,
		},
	}, nil
}
