package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryListLogic {
	return &GetDictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryListLogic) GetDictionaryList(in *core.DictionaryPageReq) (*core.DictionaryList, error) {
	var predicates []predicate.Dictionary

	if in.Name != "" {
		predicates = append(predicates, dictionary.NameContains(in.Name))
	}

	if in.Title != "" {
		predicates = append(predicates, dictionary.TitleContains(in.Title))
	}

	dicts, err := l.svcCtx.DB.Dictionary.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.DictionaryList{}
	resp.Total = dicts.PageDetails.Total

	for _, v := range dicts.List {
		resp.Data = append(resp.Data, &core.DictionaryInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			Name:      v.Name,
			Title:     v.Title,
			Status:    uint64(v.Status),
			Desc:      v.Desc,
		})
	}

	return resp, nil
}
