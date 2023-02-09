package user

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetUserListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetUserList(l.ctx, &core.UserListReq{
		Page:         req.Page,
		PageSize:     req.PageSize,
		Username:     req.Username,
		Nickname:     req.Nickname,
		Email:        req.Email,
		Mobile:       req.Mobile,
		RoleId:       req.RoleId,
		DepartmentId: req.DepartmentId,
		PositionId:   req.PositionId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserListResp{}
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.UserInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Username:     v.Username,
			Nickname:     v.Nickname,
			Mobile:       v.Mobile,
			RoleId:       v.RoleId,
			Email:        v.Email,
			Avatar:       v.Avatar,
			Status:       v.Status,
			Description:  v.Description,
			HomePath:     v.HomePath,
			DepartmentId: v.DepartmentId,
			PositionId:   v.PositionId,
		})
	}
	resp.Data.Total = data.Total
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	return resp, nil
}
