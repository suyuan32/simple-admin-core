package dictionary

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryByIdLogic {
	return &GetDictionaryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryByIdLogic) GetDictionaryById(in *core.IDReq) (*core.DictionaryInfo, error) {
	result, err := l.svcCtx.DB.Dictionary.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.DictionaryInfo{
		Id:        &result.ID,
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:    pointy.GetPointer(uint32(result.Status)),
		Title:     &result.Title,
		Name:      &result.Name,
		Desc:      &result.Desc,
		IsPublic:  &result.IsPublic,
	}, nil
}
