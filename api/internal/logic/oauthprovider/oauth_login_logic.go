package oauthprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLoginLogic) OauthLogin(req *types.OauthLoginReq) (resp *types.RedirectResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthLogin(l.ctx, &core.OauthLoginReq{
		State:    req.State,
		Provider: req.Provider,
	})
	if err != nil {
		return nil, err
	}

	return &types.RedirectResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data:         types.RedirectInfo{URL: result.Url},
	}, nil
}
