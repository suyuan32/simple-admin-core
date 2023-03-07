package dictionarydetail

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictionaryDetailLogic {
	return &DeleteDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDictionaryDetailLogic) DeleteDictionaryDetail(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.DeleteDictionaryDetail(l.ctx, &core.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
