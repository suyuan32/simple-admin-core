package role

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

type SetRoleStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleStatusLogic {
	return &SetRoleStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleStatusLogic) SetRoleStatus(req *types.SetStatusReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.SetRoleStatus(l.ctx, &core.SetStatusReq{
		Id:     req.ID,
		Status: uint64(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return &types.SimpleMsg{Msg: data.Msg}, nil
}
