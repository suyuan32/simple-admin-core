package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetApiList(context.Background(),
		&core.ApiPageReq{
			Page: &core.PageInfoReq{
				Page:     req.Page,
				PageSize: req.PageSize,
			},
			Path:        req.Path,
			Description: req.Description,
			Method:      req.Method,
			Group:       req.Group,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ApiListResp{}
	resp.Total = data.GetTotal()
	for _, v := range data.Data {
		resp.Data = append(resp.Data,
			types.ApiInfo{
				Id:          v.Id,
				CreateAt:    v.CreateAt,
				Path:        v.Path,
				Description: v.Description,
				Group:       v.Group,
				Method:      v.Method,
			})
	}
	return resp, nil
}
