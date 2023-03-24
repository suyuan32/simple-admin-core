package department

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetDepartmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepartmentListLogic) GetDepartmentList(req *types.DepartmentListReq) (resp *types.DepartmentListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetDepartmentList(l.ctx,
		&core.DepartmentListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
			Leader:   req.Leader,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.DepartmentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.DepartmentInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:     l.svcCtx.Trans.Trans(l.ctx, v.Name),
				Status:    v.Status,
				Sort:      v.Sort,
				Name:      v.Name,
				Ancestors: v.Ancestors,
				Leader:    v.Leader,
				Phone:     v.Phone,
				Email:     v.Email,
				Remark:    v.Remark,
				ParentId:  v.ParentId,
			})
	}
	return resp, nil
}
