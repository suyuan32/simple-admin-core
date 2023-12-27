package dictionarydetail

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByIdLogic {
	return &GetDictionaryDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailByIdLogic) GetDictionaryDetailById(in *core.IDReq) (*core.DictionaryDetailInfo, error) {
	result, err := l.svcCtx.DB.DictionaryDetail.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.DictionaryDetailInfo{
		Id:           &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:       pointy.GetPointer(uint32(result.Status)),
		Title:        &result.Title,
		Key:          &result.Key,
		Value:        &result.Value,
		Sort:         &result.Sort,
		DictionaryId: &result.DictionaryID,
	}, nil
}
