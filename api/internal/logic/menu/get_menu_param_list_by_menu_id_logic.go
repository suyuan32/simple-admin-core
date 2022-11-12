package menu

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParamListByMenuIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMenuParamListByMenuIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMenuParamListByMenuIdLogic {
	return &GetMenuParamListByMenuIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMenuParamListByMenuIdLogic) GetMenuParamListByMenuId(req *types.IDReq) (resp *types.MenuParamListByMenuIdResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetMenuParamListByMenuId(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	resp = &types.MenuParamListByMenuIdResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = result.Total
	for _, v := range result.Data {
		resp.Data.Data = append(resp.Data.Data, types.MenuParamInfo{
			BaseInfo: types.BaseInfo{Id: v.Id, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt},
			DataType: v.Type,
			Key:      v.Key,
			Value:    v.Value,
		})
	}

	return resp, nil
}
