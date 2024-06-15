package authority

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiAuthorityLogic {
	return &GetApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiAuthorityLogic) GetApiAuthority(req *types.IDReq) (resp *types.ApiAuthorityListResp, err error) {
	roleData, err := l.svcCtx.CoreRpc.GetRoleById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.Casbin.GetFilteredPolicy(0, *roleData.Code)
	if err != nil {
		logx.Error("failed to get old Casbin policy", logx.Field("detail", err))
		return nil, errorx.NewInternalError(err.Error())
	}

	resp = &types.ApiAuthorityListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = uint64(len(data))
	for _, v := range data {
		resp.Data.Data = append(resp.Data.Data, types.ApiAuthorityInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return resp, nil
}
