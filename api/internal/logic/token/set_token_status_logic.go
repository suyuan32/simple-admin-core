package token

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTokenStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTokenStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTokenStatusLogic {
	return &SetTokenStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetTokenStatusLogic) SetTokenStatus(req *types.SetBooleanStatusReq) (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.SetTokenStatus(context.Background(), &core.SetStatusReq{
		Id:     req.Id,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
