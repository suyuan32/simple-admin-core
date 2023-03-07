package department

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(req *types.DepartmentInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateDepartment(l.ctx,
		&core.DepartmentInfo{
			Id:        req.Id,
			Status:    req.Status,
			Sort:      req.Sort,
			Name:      req.Name,
			Ancestors: req.Ancestors,
			Leader:    req.Leader,
			Phone:     req.Phone,
			Email:     req.Email,
			Remark:    req.Remark,
			ParentId:  req.ParentId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
