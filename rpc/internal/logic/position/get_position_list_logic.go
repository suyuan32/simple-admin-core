package position

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/position"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionListLogic) GetPositionList(in *core.PositionListReq) (*core.PositionListResp, error) {
	var predicates []predicate.Position
	if in.Name != "" {
		predicates = append(predicates, position.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.Position.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.PositionListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.PositionInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			Status:    uint32(v.Status),
			Sort:      v.Sort,
			Name:      v.Name,
			Code:      v.Code,
			Remark:    v.Remark,
		})
	}

	return resp, nil
}
