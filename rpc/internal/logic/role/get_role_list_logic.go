package role

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
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

func (l *GetRoleListLogic) GetRoleList(in *core.RoleListReq) (*core.RoleListResp, error) {
	var predicates []predicate.Role
	if in.Name != "" {
		predicates = append(predicates, role.NameContains(in.Name))
	}
	if in.Code != "" {
		predicates = append(predicates, role.CodeEQ(in.Code))
	}
	if in.DefaultRouter != "" {
		predicates = append(predicates, role.DefaultRouterContains(in.DefaultRouter))
	}
	result, err := l.svcCtx.DB.Role.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize, func(pager *ent.RolePager) {
		pager.Order = ent.Asc(role.FieldSort)
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.RoleListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.RoleInfo{
			Id:            v.ID,
			CreatedAt:     v.CreatedAt.UnixMilli(),
			UpdatedAt:     v.UpdatedAt.UnixMilli(),
			Status:        uint32(v.Status),
			Name:          v.Name,
			Code:          v.Code,
			DefaultRouter: v.DefaultRouter,
			Remark:        v.Remark,
			Sort:          v.Sort,
		})
	}

	return resp, nil
}
