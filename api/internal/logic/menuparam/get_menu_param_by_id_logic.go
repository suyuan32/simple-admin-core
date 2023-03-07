package menuparam

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetMenuParamByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuParamByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamByIdLogic {
	return &GetMenuParamByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuParamByIdLogic) GetMenuParamById(req *types.IDReq) (resp *types.MenuParamInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMenuParamById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MenuParamInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.MenuParamInfo{
			BaseInfo: types.BaseInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Type:  data.Type,
			Key:   data.Key,
			Value: data.Value,
		},
	}, nil
}
