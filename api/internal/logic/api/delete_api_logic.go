package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiLogic) DeleteApi(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.DeleteApi(l.ctx, &core.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
