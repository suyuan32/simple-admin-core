package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParamListByMenuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuParamListByMenuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParamListByMenuIdLogic {
	return &GetMenuParamListByMenuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuParamListByMenuIdLogic) GetMenuParamListByMenuId(in *core.IDReq) (*core.MenuParamListResp, error) {
	params, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(in.Id)).QueryParam().All(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
		default:
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	resp := &core.MenuParamListResp{}
	resp.Total = uint64(len(params))
	for _, v := range params {
		resp.Data = append(resp.Data, &core.MenuParamResp{
			Id:        v.ID,
			Type:      v.Type,
			Key:       v.Key,
			Value:     v.Value,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
		})
	}

	return resp, nil
}
