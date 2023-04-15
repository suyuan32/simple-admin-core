package dictionarydetail

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailByDictionaryNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictionaryDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByDictionaryNameLogic {
	return &GetDictionaryDetailByDictionaryNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictionaryDetailByDictionaryNameLogic) GetDictionaryDetailByDictionaryName(req *types.DictionaryNameReq) (resp *types.DictionaryDetailListResp, err error) {
	dict, err := dictionary.NewGetDictionaryListLogic(l.ctx, l.svcCtx).GetDictionaryList(&types.DictionaryListReq{
		PageInfo: types.PageInfo{
			Page:     1,
			PageSize: 1,
		},
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	if dict.Data.Total == 0 {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	dictDetails, err := NewGetDictionaryDetailListLogic(l.ctx, l.svcCtx).
		GetDictionaryDetailList(&types.DictionaryDetailListReq{
			PageInfo: types.PageInfo{
				Page:     1,
				PageSize: 500,
			},
			Key:          "",
			DictionaryId: dict.Data.Data[0].Id,
		})

	return dictDetails, nil
}
