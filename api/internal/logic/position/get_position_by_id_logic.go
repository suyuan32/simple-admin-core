package position

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetPositionByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetPositionByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetPositionByIdLogic {
	return &GetPositionByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetPositionByIdLogic) GetPositionById(req *types.IDReq) (resp *types.PositionInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetPositionById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.PositionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.PositionInfo{
			BaseInfo: types.BaseInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status: data.Status,
			Sort:   data.Sort,
			Name:   data.Name,
			Code:   data.Code,
			Remark: data.Remark,
		},
	}, nil
}
