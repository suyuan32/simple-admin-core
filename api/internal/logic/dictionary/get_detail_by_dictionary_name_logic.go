package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailByDictionaryNameLogic {
	return &GetDetailByDictionaryNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailByDictionaryNameLogic) GetDetailByDictionaryName(req *types.DictionaryDetailReq) (resp *types.DictionaryDetailListResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetDetailByDictionaryName(l.ctx, &core.DictionaryDetailReq{Name: req.Name})

	if err != nil {
		return nil, err
	}

	resp = &types.DictionaryDetailListResp{}
	resp.Total = result.Total
	for _, v := range result.Data {
		resp.Data = append(resp.Data, types.DictionaryDetailInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Title:  v.Title,
			Key:    v.Key,
			Value:  v.Value,
			Status: v.Status,
		})
	}

	return resp, nil
}
