package api

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"

	"github.com/suyuan32/simple-admin-core/rpc/ent/api"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApiListLogic) GetApiList(in *core.ApiListReq) (*core.ApiListResp, error) {
	var predicates []predicate.API
	if in.Path != nil {
		predicates = append(predicates, api.PathContains(*in.Path))
	}
	if in.Description != nil {
		predicates = append(predicates, api.DescriptionContains(*in.Description))
	}
	if in.ApiGroup != nil {
		predicates = append(predicates, api.APIGroupContains(*in.ApiGroup))
	}
	if in.Method != nil {
		predicates = append(predicates, api.Method(*in.Method))
	}
	if in.ServiceName != nil {
		predicates = append(predicates, api.ServiceNameContains(*in.ServiceName))
	}
	result, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.ApiListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.ApiInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			Path:        &v.Path,
			Description: &v.Description,
			ApiGroup:    &v.APIGroup,
			Method:      &v.Method,
			IsRequired:  &v.IsRequired,
			ServiceName: &v.ServiceName,
		})
	}

	return resp, nil
}
