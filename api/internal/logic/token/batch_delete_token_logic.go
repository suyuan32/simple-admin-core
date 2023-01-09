package token

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewBatchDeleteTokenLogic(r *http.Request, svcCtx *svc.ServiceContext) *BatchDeleteTokenLogic {
	return &BatchDeleteTokenLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *BatchDeleteTokenLogic) BatchDeleteToken(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.BatchDeleteToken(l.ctx, &core.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
