package oauthprovider

import (
	"context"

	"github.com/chimerakang/simple-admin-core/rpc/ent/oauthprovider"

	"github.com/chimerakang/simple-admin-core/rpc/internal/svc"
	"github.com/chimerakang/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/chimerakang/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/chimerakang/simple-admin-common/i18n"
)

type DeleteOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.OauthProvider.Delete().Where(oauthprovider.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
