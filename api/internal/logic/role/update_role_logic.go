package role

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateRole(l.ctx,
		&core.RoleInfo{
			Id:     req.Id,
			Status: req.Status,
			Name:   req.Name,
			Code:   req.Code,
			Remark: req.Remark,
			Sort:   req.Sort,
		})
	if err != nil {
		return nil, err
	}

	if req.Status != nil && uint8(*req.Status) == common.StatusBanned {
		roleData, err := l.svcCtx.CoreRpc.GetRoleById(l.ctx, &core.IDReq{Id: *req.Id})
		if err != nil {
			return nil, err
		}

		// clear old policies
		var oldPolicies [][]string
		oldPolicies, err = l.svcCtx.Casbin.GetFilteredPolicy(0, *roleData.Code)
		if err != nil {
			logx.Error("failed to get old Casbin policy", logx.Field("detail", err))
			return nil, errorx.NewInternalError(err.Error())
		}

		if len(oldPolicies) != 0 {
			removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, *roleData.Code)
			if err != nil {
				l.Logger.Errorw("failed to remove roles policy", logx.Field("roleCode", roleData.Code), logx.Field("detail", err.Error()))
				return nil, errorx.NewInvalidArgumentError(err.Error())
			}
			if !removeResult {
				return nil, errorx.NewInvalidArgumentError("casbin.removeFailed")
			}
		}
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
