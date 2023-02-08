package menuparam

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/menuparam"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetMenuParamListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuParamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamListLogic {
	return &GetMenuParamListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuParamListLogic) GetMenuParamList(in *core.MenuParamListReq) (*core.MenuParamListResp, error) {
	var predicates []predicate.MenuParam
	if in.MenuId != 0 {
		predicates = append(predicates, menuparam.MenuIDEQ(in.MenuId))
	}
	result, err := l.svcCtx.DB.MenuParam.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.MenuParamListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MenuParamInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			Type:      v.Type,
			Key:       v.Key,
			Value:     v.Value,
		})
	}

	return resp, nil
}
