package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(req *types.DictionaryListReq) (resp *types.DictionaryListResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetDictionaryList(l.ctx, &core.DictionaryPageReq{
		Title:    req.Title,
		Name:     req.Name,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.DictionaryListResp{}
	resp.Total = result.Total
	for _, v := range result.Data {
		resp.Data = append(resp.Data, types.DictionaryInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Title:       v.Title,
			Name:        v.Name,
			Status:      v.Status,
			Description: v.Desc,
		})
	}

	return resp, nil
}
