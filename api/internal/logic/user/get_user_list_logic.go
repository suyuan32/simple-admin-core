package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.UserListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetUserList(l.ctx, &core.GetUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		RoleId:   req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	var res []types.UserInfoResp
	for _, v := range data.Data {
		res = append(res, types.UserInfoResp{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Username: v.Username,
			Nickname: v.Nickname,
			Mobile:   v.Mobile,
			RoleId:   v.RoleId,
			Email:    v.Email,
			Avatar:   v.Avatar,
			Status:   v.Status,
			UUID:     v.Uuid,
		})
	}
	resp = &types.UserListResp{}
	resp.Total = uint64(data.Total)
	resp.Data = res
	return resp, nil
}
