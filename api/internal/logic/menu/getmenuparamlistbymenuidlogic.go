package menu

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParamListByMenuIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuParamListByMenuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamListByMenuIdLogic {
	return &GetMenuParamListByMenuIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuParamListByMenuIdLogic) GetMenuParamListByMenuId(req *types.IDReq) (resp *types.MenuParamListByMenuIdResp, err error) {
	result, err := l.svcCtx.CoreRpc.GeMenuParamListByMenuId(l.ctx, &core.IDReq{ID: uint64(req.ID)})
	if err != nil {
		return nil, err
	}

	resp = &types.MenuParamListByMenuIdResp{}
	resp.Total = result.Total
	for _, v := range result.Data {
		resp.Data = append(resp.Data, types.MenuParamResp{
			BaseInfo: types.BaseInfo{ID: uint(v.Id), CreatedAt: v.CreateAt, UpdatedAt: v.UpdateAt},
			MenuID:   uint32(v.MenuId),
			DataType: v.Type,
			Key:      v.Key,
			Value:    v.Value,
		})
	}

	return resp, nil
}
