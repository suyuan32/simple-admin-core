package menu

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.IdReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.DeleteMenu(context.Background(), &core.IDReq{ID: uint64(req.ID)})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: result.Msg}, nil
}
