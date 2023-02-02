package department

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateDepartmentLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateDepartmentLogic {
	return &CreateOrUpdateDepartmentLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateDepartmentLogic) CreateOrUpdateDepartment(req *types.CreateOrUpdateDepartmentReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateOrUpdateDepartment(l.ctx,
		&core.DepartmentInfo{
			Id:        req.Id,
			Status:    req.Status,
			Name:      req.Name,
			Ancestors: req.Ancestors,
			Leader:    req.Leader,
			Phone:     req.Phone,
			Email:     req.Email,
			Sort:      req.Sort,
			Remark:    req.Remark,
			ParentId:  req.ParentId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
