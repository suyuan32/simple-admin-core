package menuparam

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetMenuParamListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuParamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamListLogic {
	return &GetMenuParamListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuParamListLogic) GetMenuParamList(req *types.MenuParamListReq) (resp *types.MenuParamListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuParamList(l.ctx,
		&core.MenuParamListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			MenuId:   req.MenuId,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.MenuParamListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.MenuParamInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Type:   v.Type,
				Key:    v.Key,
				Value:  v.Value,
				MenuId: v.MenuId,
			})
	}
	return resp, nil
}
