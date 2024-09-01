package configuration

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationByIdLogic {
	return &GetConfigurationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigurationByIdLogic) GetConfigurationById(in *core.IDReq) (*core.ConfigurationInfo, error) {
	result, err := l.svcCtx.DB.Configuration.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.ConfigurationInfo{
		Id:        &result.ID,
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Sort:      &result.Sort,
		State:     &result.State,
		Name:      &result.Name,
		Key:       &result.Key,
		Value:     &result.Value,
		Category:  &result.Category,
		Remark:    &result.Remark,
	}, nil
}
