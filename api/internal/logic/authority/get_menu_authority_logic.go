package authority

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

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

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.IDReq) (resp *types.MenuAuthorityInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuAuthority(l.ctx, &core.IDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.MenuAuthorityInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.MenuAuthorityInfoReq{
			RoleId:  req.Id,
			MenuIds: data.MenuId,
		},
	}

	return resp, nil
}
