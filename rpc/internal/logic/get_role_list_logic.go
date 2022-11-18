package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *core.PageInfoReq) (*core.RoleListResp, error) {
	roles, err := l.svcCtx.DB.Role.Query().Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.RoleListResp{}
	resp.Total = roles.PageDetails.Total

	for _, v := range roles.List {
		resp.Data = append(resp.Data, &core.RoleInfo{
			Id:            v.ID,
			Name:          v.Name,
			Value:         v.Value,
			DefaultRouter: v.DefaultRouter,
			Status:        uint64(v.Status),
			Remark:        v.Remark,
			OrderNo:       v.OrderNo,
			CreatedAt:     v.CreatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
