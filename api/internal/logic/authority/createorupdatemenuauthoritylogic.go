package authority

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(req *types.MenuAuthorityInfoReq) (resp *types.SimpleMsg, err error) {
	authority, err := l.svcCtx.CoreRpc.CreateOrUpdateMenuAuthority(l.ctx, &core.RoleMenuAuthorityReq{
		RoleId: req.RoleId,
		MenuId: req.MenuIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.SimpleMsg{Msg: authority.Msg}, nil
}
