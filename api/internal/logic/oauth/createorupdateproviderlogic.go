package oauth

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateProviderLogic {
	return &CreateOrUpdateProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateProviderLogic) CreateOrUpdateProvider(req *types.CreateOrUpdateProviderReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateProvider(context.Background(),
		&core.ProviderInfo{
			Id:           req.Id,
			Name:         req.Name,
			ClientId:     req.ClientID,
			ClientSecret: req.ClientSecret,
			RedirectUrl:  req.RedirectURL,
			Scopes:       req.Scopes,
			AuthUrl:      req.AuthURL,
			TokenUrl:     req.TokenURL,
			AuthStyle:    uint64(req.AuthStyle),
			InfoUrl:      req.InfoURL,
			CreateAt:     req.CreateAt,
		})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
