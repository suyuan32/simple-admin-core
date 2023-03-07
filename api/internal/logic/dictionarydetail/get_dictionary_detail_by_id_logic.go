package dictionarydetail

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type GetDictionaryDetailByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetDictionaryDetailByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetDictionaryDetailByIdLogic {
	return &GetDictionaryDetailByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetDictionaryDetailByIdLogic) GetDictionaryDetailById(req *types.IDReq) (resp *types.DictionaryDetailInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetDictionaryDetailById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DictionaryDetailInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.DictionaryDetailInfo{
			BaseInfo: types.BaseInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:       data.Status,
			Title:        data.Title,
			Key:          data.Key,
			Value:        data.Value,
			DictionaryId: data.DictionaryId,
			Sort:         data.Sort,
		},
	}, nil
}
