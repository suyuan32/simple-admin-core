package authority

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.IdReq) (resp *types.MenuAuthorityInfo, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuAuthority(context.Background(), &core.IDReq{
		ID: uint64(req.ID),
	})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuAuthorityInfo{
		RoleId:  uint64(req.ID),
		MenuIds: []uint64{},
	}

	for _, v := range data.MenuId {
		resp.MenuIds = append(resp.MenuIds, v)
	}
	return resp, nil
}
