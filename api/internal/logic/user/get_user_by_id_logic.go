package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.UUIDReq) (resp *types.UserInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UserInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:       data.Status,
			Username:     data.Username,
			Nickname:     data.Nickname,
			Description:  data.Description,
			HomePath:     data.HomePath,
			RoleIds:      data.RoleIds,
			Mobile:       data.Mobile,
			Email:        data.Email,
			Avatar:       data.Avatar,
			DepartmentId: data.DepartmentId,
			PositionIds:  data.PositionIds,
		},
	}, nil
}
