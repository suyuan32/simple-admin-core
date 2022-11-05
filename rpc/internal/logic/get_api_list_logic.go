package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/api"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var predicates []predicate.Api

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

	pageResult, err := l.svcCtx.DB.Api.Query().Where(predicates...).Page(l.ctx, req.PageNo, req.PageSize)

	if result.Error != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", result.Error.Error()))
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	for _, v := range apis {
		resp.Data = append(resp.Data, &core.ApiInfo{
			Id:          uint64(v.ID),
			CreatedAt:   v.CreatedAt.UnixMilli(),
			Path:        v.Path,
			Description: v.Description,
			Group:       v.ApiGroup,
			Method:      v.Method,
		})
	}
	return resp, nil
}
