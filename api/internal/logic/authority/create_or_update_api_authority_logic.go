package authority

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiAuthorityLogic {
	return &CreateOrUpdateApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateApiAuthorityLogic) CreateOrUpdateApiAuthority(req *types.CreateOrUpdateApiAuthorityReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetRoleById(l.ctx, &core.IDReq{Id: req.RoleId})
	if err != nil {
		return nil, err
	}

	// clear old policies
	var oldPolicies [][]string
	oldPolicies = l.svcCtx.Casbin.GetFilteredPolicy(0, data.Code)
	if len(oldPolicies) != 0 {
		removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, data.Code)
		if err != nil {
			return &types.BaseMsgResp{
				Code: errorcode.Internal,
				Msg:  err.Error(),
			}, nil
		}
		if !removeResult {
			return &types.BaseMsgResp{
				Code: errorcode.Internal,
				Msg:  l.svcCtx.Trans.Trans(l.ctx, "casbin.removeFailed"),
			}, nil
		}
	}
	// add new policies
	var policies [][]string
	for _, v := range req.Data {
		policies = append(policies, []string{data.Code, v.Path, v.Method})
	}
	addResult, err := l.svcCtx.Casbin.AddPolicies(policies)
	if err != nil {
		return &types.BaseMsgResp{
			Code: errorcode.Internal,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, "casbin.addFailed"),
		}, nil
	}
	if addResult {
		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
	} else {
		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateFailed)}, nil
	}
}
