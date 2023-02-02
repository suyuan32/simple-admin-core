package department

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type GetDepartmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetDepartmentListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
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
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.DepartmentInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
			})
	}
	return resp, nil
}
