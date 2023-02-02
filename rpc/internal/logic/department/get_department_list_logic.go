package department

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/department"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetDepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentListLogic) GetDepartmentList(in *core.DepartmentPageReq) (*core.DepartmentListResp, error) {
	var predicates []predicate.Department
	if in.Name != "" {
		predicates = append(predicates, department.NameContains(in.Name))
	}
	if in.Ancestors != "" {
		predicates = append(predicates, department.AncestorsContains(in.Ancestors))
	}
	if in.Leader != "" {
		predicates = append(predicates, department.LeaderContains(in.Leader))
	}
	result, err := l.svcCtx.DB.Department.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.DepartmentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.DepartmentInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			Status:    uint64(v.Status),
			Name:      v.Name,
			Ancestors: v.Ancestors,
			Leader:    v.Leader,
			Phone:     v.Phone,
			Email:     v.Email,
			Sort:      v.Sort,
			ParentId:  v.ParentID,
		})
	}

	return resp, nil
}
