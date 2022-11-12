package oauth

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateProviderLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateProviderLogic {
	return &CreateOrUpdateProviderLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateProviderLogic) CreateOrUpdateProvider(req *types.CreateOrUpdateProviderReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.CreateOrUpdateProvider(l.ctx,
		&core.ProviderInfo{
			Id:           req.Id,
			Name:         req.Name,
			ClientId:     req.ClientId,
			ClientSecret: req.ClientSecret,
			RedirectUrl:  req.RedirectURL,
			Scopes:       req.Scopes,
			AuthUrl:      req.AuthURL,
			TokenUrl:     req.TokenURL,
			AuthStyle:    uint64(req.AuthStyle),
			InfoUrl:      req.InfoURL,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
