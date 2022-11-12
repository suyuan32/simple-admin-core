package dictionary

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetDetailByDictionaryNameLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetDetailByDictionaryNameLogic {
	return &GetDetailByDictionaryNameLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetDetailByDictionaryNameLogic) GetDetailByDictionaryName(req *types.DictionaryDetailReq) (resp *types.DictionaryDetailListResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetDetailByDictionaryName(l.ctx, &core.DictionaryDetailReq{Name: req.Name})

	if err != nil {
		return nil, err
	}

	resp = &types.DictionaryDetailListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = result.Total
	for _, v := range result.Data {
		resp.Data.Data = append(resp.Data.Data, types.DictionaryDetailInfo{
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
