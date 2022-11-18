package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/api"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
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

func (l *GetApiListLogic) GetApiList(in *core.ApiPageReq) (*core.ApiListResp, error) {
	var predicates []predicate.API

	if in.Path != "" {
		predicates = append(predicates, api.PathContains(in.Path))
	}

	if in.Description != "" {
		predicates = append(predicates, api.DescriptionContains(in.Description))
	}

	if in.Method != "" {
		predicates = append(predicates, api.MethodContains(in.Method))
	}

	if in.Group != "" {
		predicates = append(predicates, api.APIGroupContains(in.Group))
	}

	apis, err := l.svcCtx.DB.API.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.ApiListResp{}
	resp.Total = apis.PageDetails.Total

	for _, v := range apis.List {
		resp.Data = append(resp.Data, &core.ApiInfo{
			Id:          v.ID,
			CreatedAt:   v.CreatedAt.UnixMilli(),
			Path:        v.Path,
			Description: v.Description,
			Group:       v.APIGroup,
			Method:      v.Method,
		})
	}

	return resp, nil
}
