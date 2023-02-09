package dictionarydetail

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type GetDictionaryDetailListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetDictionaryDetailListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(req *types.DictionaryDetailListReq) (resp *types.DictionaryDetailListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetDictionaryDetailList(l.ctx,
		&core.DictionaryDetailListReq{
			Page:         req.Page,
			PageSize:     req.PageSize,
			DictionaryId: req.DictionaryId,
			Key:          req.Key,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.DictionaryDetailListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.DictionaryDetailInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:       v.Status,
				Title:        v.Title,
				Key:          v.Key,
				Value:        v.Value,
				DictionaryId: v.DictionaryId,
				Sort:         v.Sort,
			})
	}
	return resp, nil
}
