package dictionarydetail

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

type GetDictionaryDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailListLogic {
	return &GetDictionaryDetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailListLogic) GetDictionaryDetailList(in *core.DictionaryDetailListReq) (*core.DictionaryDetailListResp, error) {
	var predicates []predicate.DictionaryDetail
	if in.DictionaryId != 0 {
		predicates = append(predicates, dictionarydetail.DictionaryIDEQ(in.DictionaryId))
	}
	if in.Key != "" {
		predicates = append(predicates, dictionarydetail.KeyContains(in.Key))
	}
	result, err := l.svcCtx.DB.DictionaryDetail.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize, func(pager *ent.DictionaryDetailPager) {
		pager.Order = ent.Asc(dictionarydetail.FieldSort)
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.DictionaryDetailListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.DictionaryDetailInfo{
			Id:           v.ID,
			CreatedAt:    v.CreatedAt.UnixMilli(),
			UpdatedAt:    v.UpdatedAt.UnixMilli(),
			Status:       uint32(v.Status),
			Title:        v.Title,
			Key:          v.Key,
			Value:        v.Value,
			DictionaryId: v.DictionaryID,
			Sort:         v.Sort,
		})
	}

	return resp, nil
}
