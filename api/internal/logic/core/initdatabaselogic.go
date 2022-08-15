package core

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.SimpleMsg, err error) {
	result, err := l.svcCtx.CoreRpc.InitDatabase(context.Background(), &core.Empty{})
	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: result.Msg}, nil
}
