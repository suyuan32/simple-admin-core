package api

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewDeleteApiLogic(r *http.Request, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *DeleteApiLogic) DeleteApi(req *types.IDReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.DeleteApi(l.ctx, &core.IDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.SimpleMsg{}
	resp.Msg = data.Msg
	return resp, nil
}
