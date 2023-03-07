package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOauthProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOauthProviderLogic {
	return &UpdateOauthProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOauthProviderLogic) UpdateOauthProvider(req *types.OauthProviderInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateOauthProvider(l.ctx,
		&core.OauthProviderInfo{
			Id:           req.Id,
			Name:         req.Name,
			ClientId:     req.ClientId,
			ClientSecret: req.ClientSecret,
			RedirectUrl:  req.RedirectUrl,
			Scopes:       req.Scopes,
			AuthUrl:      req.AuthUrl,
			TokenUrl:     req.TokenUrl,
			AuthStyle:    req.AuthStyle,
			InfoUrl:      req.InfoUrl,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
