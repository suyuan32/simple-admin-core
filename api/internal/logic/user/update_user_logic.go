package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateUser(l.ctx,
		&core.UserInfo{
			Id:           req.Id,
			Status:       req.Status,
			Username:     req.Username,
			Password:     req.Password,
			Nickname:     req.Nickname,
			Description:  req.Description,
			HomePath:     req.HomePath,
			RoleIds:      req.RoleIds,
			Mobile:       req.Mobile,
			Email:        req.Email,
			Avatar:       req.Avatar,
			DepartmentId: req.DepartmentId,
			PositionIds:  req.PositionIds,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
