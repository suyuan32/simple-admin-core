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

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.UserListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetUserList(l.ctx, &core.GetUserListReq{
		Page:         req.Page,
		PageSize:     req.PageSize,
		Username:     req.Username,
		Nickname:     req.Nickname,
		Email:        req.Email,
		Mobile:       req.Mobile,
		RoleId:       req.RoleId,
		DepartmentId: req.DepartmentId,
	})
	if err != nil {
		return nil, err
	}
	var res []types.UserInfoResp
	for _, v := range data.Data {
		res = append(res, types.UserInfoResp{
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
		})
	}
	resp = &types.UserListResp{}
	resp.Data.Total = data.Total
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Data = res
	return resp, nil
}
