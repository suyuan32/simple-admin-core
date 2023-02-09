package api

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetApiByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetApiByIdLogic) GetApiById(req *types.IDReq) (resp *types.ApiInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetApiById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.ApiInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.ApiInfo{
			BaseInfo:    types.BaseInfo{Id: data.Id, CreatedAt: data.CreatedAt, UpdatedAt: data.UpdatedAt},
			Trans:       l.svcCtx.Trans.Trans(l.lang, data.Description),
			Path:        data.Path,
			Description: data.Description,
			Group:       data.ApiGroup,
			Method:      data.Method,
		},
	}, nil
}
