package role

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetRoleListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.PageInfo) (resp *types.RoleListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetRoleList(l.ctx, &core.PageInfoReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.RoleListResp{}
	resp.Total = data.Total

	lang := l.r.Header.Get("Accept-Language")

	for _, v := range data.Data {
		resp.Data = append(resp.Data, types.RoleInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
			},
			Title:         l.svcCtx.Trans.Trans(lang, v.Name),
			Name:          v.Name,
			Value:         v.Value,
			DefaultRouter: v.DefaultRouter,
			Status:        uint32(v.Status),
			Remark:        v.Remark,
			OrderNo:       v.OrderNo,
		})
	}
	return resp, nil
}
